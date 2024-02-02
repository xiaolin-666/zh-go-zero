package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Follow struct {
	ID             int64 `gorm:"primary_key"`
	UserID         int64
	FollowedUserID int64
	FollowStatus   int
	CreateTime     time.Time
	UpdateTime     time.Time
}

func (f *Follow) TableName() string {
	return "follow"
}

type FollowModel struct {
	db *gorm.DB
}

func NewFollowModel(db *gorm.DB) *FollowModel {
	return &FollowModel{db: db}
}

func (followModel *FollowModel) FindFollowByUserIDAndFollowUserID(ctx context.Context, userID int64, followUserID int64) (*Follow, error) {
	var follow Follow

	err := followModel.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Where("followed_user_id = ?", followUserID).Find(&follow).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &follow, err
}

func (followModel *FollowModel) Insert(ctx context.Context, data *Follow) error {
	return followModel.db.WithContext(ctx).Create(data).Error
}

func (followModel *FollowModel) Update(ctx context.Context, data *Follow) error {
	return followModel.db.WithContext(ctx).Save(data).Error
}

func (followModel *FollowModel) UpdateFields(ctx context.Context, id int64, values map[string]interface{}) error {
	return followModel.db.WithContext(ctx).Model(&Follow{}).Where("id = ?", id).Updates(values).Error
}
