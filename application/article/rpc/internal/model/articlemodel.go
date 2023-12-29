package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"time"
	"zh-go-zero/application/article/rpc/internal/types"
)

var _ ArticleModel = (*customArticleModel)(nil)

type (
	// ArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleModel.
	ArticleModel interface {
		articleModel
		ArticlesByAuthorId(ctx context.Context, userId int64, sortType int32, cursor int64, limit int) ([]*Article, error)
	}

	customArticleModel struct {
		*defaultArticleModel
	}
)

// NewArticleModel returns a model for the database table.
func NewArticleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ArticleModel {
	return &customArticleModel{
		defaultArticleModel: newArticleModel(conn, c, opts...),
	}
}

func (m *defaultArticleModel) ArticlesByAuthorId(ctx context.Context, userId int64, sortType int32, cursor int64, limit int) ([]*Article, error) {
	var (
		sql      string
		resp     []*Article
		anyField string
	)
	if sortType == types.SortPublishTime {
		anyField = time.Unix(cursor, 0).Format("2006-01-02 15:04:05")
		sql = fmt.Sprint("select " + articleRows + " from " + m.table + " where `author_id` = ? and `publish_time` < ? order by `publish_time` desc limit ?")
	} else {
		anyField = strconv.Itoa(types.DefaultSortLikeCursor)
		sql = fmt.Sprint("select " + articleRows + " from " + m.table + " where `author_id` = ? and `like_num` < ? order by `like_num` desc limit ?")
	}
	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, userId, anyField, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
