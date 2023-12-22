package types

type ThumbUpMsg struct {
	BizId    string `json:"bizId,omitempty"`    // 业务id
	ObjId    int64  `json:"objId,omitempty"`    // 对象id
	UserId   int64  `json:"userId,omitempty"`   // 用户id
	LikeType int32  `json:"likeType,omitempty"` // 点赞类型
}
