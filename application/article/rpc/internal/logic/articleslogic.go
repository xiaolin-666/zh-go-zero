package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
	"slices"
	"strconv"
	"time"
	"zh-go-zero/application/article/rpc/internal/code"
	"zh-go-zero/application/article/rpc/internal/model"
	"zh-go-zero/application/article/rpc/internal/types"

	"zh-go-zero/application/article/rpc/internal/svc"
	"zh-go-zero/application/article/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixArticles = "biz#articles#%d#%d"
	articlesExpire = 3600 * 24 * 2
)

type ArticlesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticlesLogic {
	return &ArticlesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ArticlesLogic) Articles(in *service.ArticlesRequest) (*service.ArticlesResponse, error) {
	if in.UserId < 0 {
		return nil, code.UserIdInvalid
	}
	if in.SortType != types.SortPublishTime && in.SortType != types.SortLikeCount {
		return nil, code.SortTypeInvalid
	}
	if in.PageSize == 0 {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor == 0 {
		if in.SortType == types.SortPublishTime {
			in.Cursor = time.Now().Unix()
		} else {
			in.Cursor = types.SortLikeCount
		}
	}

	var (
		err            error
		isCache, isEnd bool
		articles       []*model.Article
		curPage        []*service.ArticleItem
	)

	articleIds, _ := l.cacheArticles(l.ctx, in.UserId, in.PageSize, in.Cursor, in.SortType)
	if len(articleIds) > 0 {
		isCache = true
		if articleIds[len(articleIds)-1] == -1 {
			isEnd = true
		}
		articles, err = l.articleByIds(l.ctx, articleIds)
		if err != nil {
			return nil, err
		}
		// mr.MapReduce结果后需排序
		var cmpFunc func(i, j *model.Article) bool
		if in.SortType == types.SortPublishTime {
			cmpFunc = func(i, j *model.Article) bool {
				return i.PublishTime.Unix() > j.PublishTime.Unix()
			}
		} else {
			cmpFunc = func(i, j *model.Article) bool {
				return i.LikeNum > j.LikeNum
			}
		}
		articles = slices.CompactFunc(articles, cmpFunc)
		for _, article := range articles {
			curPage = append(curPage, &service.ArticleItem{
				ArticleId:    article.Id,
				Title:        article.Title,
				Content:      article.Content,
				Description:  article.Description,
				Cover:        article.Cover,
				CommentCount: article.CommentNum,
				LikeCount:    article.LikeNum,
				PublishTime:  article.PublishTime.Unix(),
			})
		}
	} else {
		v, err, _ := l.svcCtx.SingleFlight.Do(fmt.Sprintf("ArticlesByUserId%v%v", in.UserId, in.SortType), func() (interface{}, error) {
			return l.svcCtx.ArticleModel.ArticlesByAuthorId(l.ctx, in.UserId, in.SortType, in.Cursor, types.DefaultLimit)
		})
		if err != nil {
			return nil, err
		}
		if v == nil {
			return &service.ArticlesResponse{}, nil
		}
		articles = v.([]*model.Article)
		var firstPage []*model.Article
		if len(articles) > types.DefaultPageSize {
			firstPage = articles[:types.DefaultPageSize]
		} else {
			firstPage = articles
			isEnd = true
		}
		for _, article := range firstPage {
			curPage = append(curPage, &service.ArticleItem{
				ArticleId:    article.Id,
				Title:        article.Title,
				Content:      article.Content,
				Description:  article.Description,
				Cover:        article.Cover,
				CommentCount: article.CommentNum,
				LikeCount:    article.LikeNum,
				PublishTime:  article.PublishTime.Unix(),
			})
		}
	}
	var (
		cursor, lastId int64
	)

	if len(curPage) > 0 {
		lastPage := curPage[len(curPage)-1]
		lastId = lastPage.ArticleId
		if in.Cursor == types.SortPublishTime {
			cursor = lastPage.PublishTime
		} else {
			cursor = lastPage.LikeCount
		}
		for i, article := range curPage {
			if in.Cursor == types.SortPublishTime {
				if article.PublishTime == in.Cursor && in.ArticleId == article.ArticleId {
					curPage = curPage[i+1:]
				}
			} else {
				if article.LikeCount == in.Cursor && in.ArticleId == article.ArticleId {
					curPage = curPage[i+1:]
				}
			}
		}
	}
	if !isCache {
		threading.GoSafe(func() {
			// 总数据不足够DefaultLimit时, 可以判断出isEnd
			if len(articles) < types.DefaultLimit && len(articles) > 0 {
				articles = append(articles, &model.Article{Id: -1})
			}
			err = l.addArticlesCache(articles, in.UserId, in.SortType)
			if err != nil {
				logx.Errorf("addArticlesCache err: %v", err)
			}
		})
	}

	return &service.ArticlesResponse{
		Articles:  curPage,
		IsEnd:     isEnd,
		Cursor:    cursor,
		ArticleId: lastId,
	}, nil
}

func (l *ArticlesLogic) addArticlesCache(articles []*model.Article, userId int64, sortType int32) error {
	key := articlesKey(userId, sortType)
	var score int64

	for _, article := range articles {
		if sortType == types.SortPublishTime && article.Id != -1 {
			score = article.PublishTime.Local().Unix()
		} else if sortType == types.SortLikeCount {
			score = article.LikeNum
		}
		_, err := l.svcCtx.BizRedis.Zadd(key, score, strconv.FormatInt(article.Id, 10))
		if err != nil {
			return err
		}
	}
	return l.svcCtx.BizRedis.Expire(key, articlesExpire)
}

func (l *ArticlesLogic) cacheArticles(ctx context.Context, userId, pageSize, cursor int64, sortType int32) ([]int64, error) {
	var articleIds []int64
	artKey := articlesKey(userId, sortType)
	exits, err := l.svcCtx.BizRedis.ExistsCtx(ctx, artKey)
	if err != nil {
		logx.Errorf("l.svcCtx.BizRedis.ExistsCtx error: %v", err)
	}
	if exits {
		err = l.svcCtx.BizRedis.ExpireCtx(ctx, artKey, articlesExpire)
		if err != nil {
			logx.Errorf("ExpireCtx key: %s error: %v", artKey, err)
		}
	}
	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, artKey, 0, cursor, 0, int(pageSize))
	if err != nil {
		logx.Errorf("ZrevrangebyscoreWithScoresAndLimit key: %s error: %v", artKey, err)
		return nil, err
	}
	for _, pair := range pairs {
		artId, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			logx.Errorf("ZrevrangebyscoreWithScoresAndLimit key: %s error: %v", artKey, err)
			return nil, err
		}
		articleIds = append(articleIds, artId)
	}
	return articleIds, nil
}

func (l *ArticlesLogic) articleByIds(ctx context.Context, articleIds []int64) ([]*model.Article, error) {
	articles, err := mr.MapReduce[int64, *model.Article, []*model.Article](func(source chan<- int64) {
		for _, aid := range articleIds {
			if aid == -1 {
				continue
			}
			source <- aid
		}
	}, func(id int64, writer mr.Writer[*model.Article], cancel func(error)) {
		article, err := l.svcCtx.ArticleModel.FindOne(ctx, id)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(article)
	}, func(pipe <-chan *model.Article, writer mr.Writer[[]*model.Article], cancel func(error)) {
		var articles []*model.Article
		for article := range pipe {
			articles = append(articles, article)
		}
		writer.Write(articles)
	})
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func articlesKey(userId int64, sortType int32) string {
	return fmt.Sprintf("%v%v%v", prefixArticles, userId, sortType)
}
