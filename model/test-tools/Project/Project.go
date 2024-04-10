package Project

import (
	"test-case-gin/utils"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          string         `gorm:"type:varchar(255)" json:"id"`          // 项目ID
	Name        string         `gorm:"type:varchar(255)" json:"name"`        // 项目名称
	Description string         `gorm:"type:varchar(255)" json:"description"` // 项目描述
	SerialNum   string         `gorm:"type:varchar(255)" json:"serialNum"`   // 项目编号
	Category    string         `gorm:"type:varchar(255)" json:"category"`    // 项目分类
	CreateBy    string         `gorm:"type:varchar(255)" json:"createBy"`    // 创建人
	CreatedAt   time.Time      `json:"created_at"`                           // 创建时间
	UpdatedAt   time.Time      `json:"updated_at"`                           // 更新时间
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// BeforeCreate 钩子函数 创建之前
func (u *Project) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = utils.RandAllString(32)
	return nil
}
