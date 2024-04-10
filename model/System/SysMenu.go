package System

import (
	"test-case-gin/utils"
	"time"

	"gorm.io/gorm"
)

type SysMenu struct {
	MenuId      string         `gorm:"type:varchar(255)" json:"menu_id" COMMENT:"菜单id"`
	MenuName    string         `gorm:"type:varchar(255)" json:"menu_name" COMMENT:"菜单名称"`
	ParentId    string         `gorm:"type:varchar(255)" json:"parent_id" COMMENT:"父菜单id"`
	OrderNum    int            `gorm:"type:int" json:"order_num" COMMENT:"显示顺序"`
	Path        string         `gorm:"type:varchar(255)" json:"path" COMMENT:"路由地址"`
	Component   string         `gorm:"type:varchar(255)" json:"component" COMMENT:"组件路径"`
	Query       string         `gorm:"type:varchar(255)" json:"query" COMMENT:"路由参数"`
	IsFrame     int            `gorm:"type:int" json:"is_frame" COMMENT:"是否为外链 0 是 1 否"`
	IsCache     int            `gorm:"type:int" json:"is_cache" COMMENT:"是否缓存 0 缓存 1 不缓存"`
	MenuType    string         `gorm:"type:varchar(255)" json:"menu_type" COMMENT:"菜单类型 0 目录 1菜单 2 按钮"`
	Visible     string         `gorm:"type:varchar(255)" json:"visible" COMMENT:"菜单状态 0 显示 1隐藏"`
	Status      string         `gorm:"type:varchar(255)" json:"status" COMMENT:"菜单状态 0 正常 1 停用"`
	Perms       string         `gorm:"type:varchar(255)" json:"perms" COMMENT:"权限标识"`
	Icon        string         `gorm:"type:varchar(255)" json:"icon" COMMENT:"菜单图标"`
	CreateBy    string         `gorm:"type:varchar(255)" json:"create_by" COMMENT:"创建者"`
	CreatedTime *time.Time     `gorm:"type:datetime;default:current_timestamp" json:"created_time"   COMMENT:"创建时间"`
	UpdatedTime *time.Time     `gorm:"type:datetime null;" json:"updated_time" COMMENT:"更新时间"` // 更新时间
	DeletedTime gorm.DeletedAt `gorm:"index" json:"deleted_time" COMMENT:"删除时间"`
	Remark      string         `gorm:"type:varchar(500)" json:"remark" COMMENT:"备注"`
}

// BeforeCreate 钩子函数 创建之前
func (u *SysMenu) BeforeCreate(_ *gorm.DB) (err error) {
	u.MenuId = utils.RandAllString(32)
	return nil
}
