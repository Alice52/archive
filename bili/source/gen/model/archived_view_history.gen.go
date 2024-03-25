// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameArchivedViewHistory = "archived_view_history"

// ArchivedViewHistory 浏览历史记录
type ArchivedViewHistory struct {
	ID         int64          `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	CreateTime *int64         `gorm:"column:create_time;type:bigint unsigned;autoCreateTime" json:"create_time"`
	UpdateTime *int64         `gorm:"column:update_time;type:bigint unsigned;autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:bigint" json:"delete_time"`
}

// TableName ArchivedViewHistory's table name
func (*ArchivedViewHistory) TableName() string {
	return TableNameArchivedViewHistory
}