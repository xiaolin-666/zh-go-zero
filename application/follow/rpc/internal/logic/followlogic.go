package logic

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
	"zh-go-zero/application/follow/rpc/code"
	"zh-go-zero/application/follow/rpc/internal/model"
	"zh-go-zero/application/follow/rpc/internal/types"

	"zh-go-zero/application/follow/rpc/internal/svc"
	"zh-go-zero/application/follow/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *service.FollowRequest) (*service.FollowResponse, error) {
	if in.UserID == 0 {
		return &service.FollowResponse{}, code.FollowUserIdEmpty
	}
	if in.FollowUserID == 0 {
		return &service.FollowResponse{}, code.FollowedUserIdEmpty
	}
	if in.UserID == in.FollowUserID {
		return &service.FollowResponse{}, code.CannotFollowSelf
	}
	follow, err := l.svcCtx.FollowModel.FindFollowByUserIDAndFollowUserID(l.ctx, in.UserID, in.FollowUserID)
	if err != nil {
		l.Logger.Errorf("[Follow] FollowModel.FindByUserIDAndFollowedUserID err: %v req: %v", err, in)
		return nil, err
	}
	if follow != nil && follow.FollowStatus == types.FollowStatusFollow {
		return nil, nil
	}

	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		if follow != nil && follow.FollowStatus == types.FollowStatusUnfollow {
			err = model.NewFollowModel(tx).UpdateFields(l.ctx, follow.ID, map[string]interface{}{"follow_status": types.FollowStatusFollow})
			if err != nil {
				return err
			}
		} else {
			err = model.NewFollowModel(tx).Insert(l.ctx, &model.Follow{
				UserID:         in.UserID,
				FollowedUserID: in.FollowUserID,
				FollowStatus:   types.FollowStatusFollow,
				CreateTime:     time.Now(),
				UpdateTime:     time.Now(),
			})
			err = model.NewFollowCountModel(tx).IncrFollowCount(l.ctx, in.UserID)
			if err != nil {
				return err
			}
			err = model.NewFollowCountModel(tx).IncrFansCount(l.ctx, in.FollowUserID)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		l.Logger.Errorf("[Follow] Transaction error: %v", err)
		return nil, err
	}

	followExist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, userFollowKey(in.UserID))
	if err != nil {
		l.Logger.Errorf("[Follow] Redis Exists error: %v", err)
		return nil, err
	}
	if followExist {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, userFollowKey(in.UserID), time.Now().Unix(), strconv.FormatInt(in.FollowUserID, 10))
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zadd error: %v", err)
			return nil, err
		}
		_, err = l.svcCtx.BizRedis.ZremrangebyrankCtx(l.ctx, userFollowKey(in.UserID), 0, types.CacheMaxFollowCount)
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zremrangebyrank error: %v", err)
		}
	}

	fansExist, err := l.svcCtx.BizRedis.ExistsCtx(l.ctx, userFansKey(in.FollowUserID))
	if err != nil {
		l.Logger.Errorf("[Follow] Redis Exists error: %v", err)
		return nil, err
	}
	if fansExist {
		_, err = l.svcCtx.BizRedis.ZaddCtx(l.ctx, userFansKey(in.FollowUserID), time.Now().Unix(), strconv.FormatInt(in.UserID, 10))
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zadd error: %v", err)
			return nil, err
		}
		_, err = l.svcCtx.BizRedis.ZremrangebyrankCtx(l.ctx, userFansKey(in.FollowUserID), 0, types.CacheMaxFansCount)
		if err != nil {
			l.Logger.Errorf("[Follow] Redis Zremrangebyrank error: %v", err)
		}
	}

	return &service.FollowResponse{}, nil
}

func userFollowKey(userId int64) string {
	return fmt.Sprintf("biz#user#follow#%d", userId)
}

func userFansKey(userId int64) string {
	return fmt.Sprintf("biz#user#fans#%d", userId)
}
