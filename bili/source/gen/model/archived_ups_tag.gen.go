// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArchivedUpsTag = "archived_ups_tag"

// ArchivedUpsTag 关注的UP主分组
type ArchivedUpsTag struct {
	TagID      int64          `gorm:"column:tag_id;type:bigint;primaryKey;comment:tagid" json:"tag_id"` // tagid
	CreateTime *time.Time     `gorm:"column:create_time;type:datetime(3);autoCreateTime" json:"create_time"`
	UpdateTime *time.Time     `gorm:"column:update_time;type:datetime(3);autoUpdateTime" json:"update_time"`
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;type:datetime(3)" json:"delete_time"`
	Name       *string        `gorm:"column:name;type:varchar(64);comment:name" json:"name"` // name
	Count_     *int64         `gorm:"column:count;type:bigint;comment:count" json:"count"`   // count
	Resp       *string        `gorm:"column:resp;type:json" json:"resp"`
	Tip        *string        `gorm:"column:tip;type:varchar(128);comment:tip" json:"tip"` // tip
	ArchivedUp []ArchivedUp   `gorm:"foreignKey:tag_id" json:"archived_up"`
}

// TableName ArchivedUpsTag's table name
func (*ArchivedUpsTag) TableName() string {
	return TableNameArchivedUpsTag
}
