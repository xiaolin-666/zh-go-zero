package model

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type FollowCount struct {
	ID          int64 `gorm:"primary_key"`
	UserID      int64
	FollowCount int64
	FansCount   int
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (followCount *FollowCount) TableName() string {
	return "follow_count"
}

type FollowCountModel struct {
	db *gorm.DB
}

func NewFollowCountModel(db *gorm.DB) *FollowCountModel {
	return &FollowCountModel{db: db}
}

func (followCountModel *FollowCountModel) IncrFollowCount(ctx context.Context, userID int64) error {
	return followCountModel.db.WithContext(ctx).
		Exec("insert into follow_count (user_id, follow_count) values (?, 1) on duplicate key update follow_count = follow_count + 1", userID).
		Error
}

func (followCountModel *FollowCountModel) IncrFansCount(ctx context.Context, userID int64) error {
	return followCountModel.db.WithContext(ctx).
		Exec("insert into follow_count (user_id, fans_count) values (?, 1) on duplicate key update fans_count = fans_count + 1", userID).
		Error
}
