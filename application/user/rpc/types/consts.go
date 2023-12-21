package types

const (
	ArticleStatusPending = iota // 待审核
	ArticleStatusNotPass = iota // 不同过
	ArticleStatusVisible = iota // 可见
	ArticleStatusDeleted = iota // 已删除
)
