package logic

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"zh-go-zero/application/article/api/code"

	"zh-go-zero/application/article/api/internal/svc"
	"zh-go-zero/application/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover(r *http.Request) (resp *types.UploadCoverResponse, err error) {
	r.Body = http.MaxBytesReader(nil, r.Body, maxFileSize)
	err = r.ParseMultipartForm(maxFileSize)
	if err != nil {
		return nil, code.CoverTooBigErr
	}
	file, header, err := r.FormFile("cover")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.BucketName)
	if err != nil {
		logx.Errorf("get bucket fialed, err: %v", err)
		return nil, code.GetBucketErr
	}
	objectKey := genFileName(header.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("upload cover fialed, err: %v", err)
		return nil, code.PutBucketErr
	}
	return &types.UploadCoverResponse{CoverUrl: genCoverUrl(objectKey)}, nil
}

func genFileName(fn string) string {
	return fmt.Sprintf("%d_%s", time.Now().UnixMilli(), fn)
}

func genCoverUrl(objectKey string) string {
	return fmt.Sprintf("https://zh-go-zero.oss-cn-beijing.aliyuncs.com/%s", objectKey)
}
