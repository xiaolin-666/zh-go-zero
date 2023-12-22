package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/threading"
	"zh-go-zero/application/like/rpc/internal/types"

	"zh-go-zero/application/like/rpc/internal/svc"
	"zh-go-zero/application/like/rpc/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbUpLogic {
	return &ThumbUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbUpLogic) ThumbUp(in *service.ThumbUpRequest) (*service.ThumbUpResponse, error) {
	msg := types.ThumbUpMsg{
		BizId:    in.BizId,
		UserId:   in.UserId,
		ObjId:    in.ObjId,
		LikeType: in.LikeType,
	}
	// TODO kq test
	threading.GoSafe(func() {
		data, err := json.Marshal(&msg)
		if err != nil {
			logx.Errorf("ThumbUp json marshal error: %v, req: %v", err, in)
			return
		}
		if err := l.svcCtx.KqPusherClient.Push(string(data)); err != nil {
			logx.Errorf("KqPusherClient Push Error , err :%v", err)
		}
	})

	return &service.ThumbUpResponse{}, nil
}
