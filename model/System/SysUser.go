package System

import (
	"gorm.io/gorm"
	"time"
)

type SysUser struct {
	UserID      int            `gorm:"primary_key;type:int" json:"user_id" COMMENT:"用户ID"`
	DeptID      string         `gorm:"type:varchar(255)" json:"dept_id" COMMENT:"部门ID"`
	CreatedTime *time.Time     `gorm:"type:datetime;default:current_timestamp" json:"created_time"   COMMENT:"创建时间"`
	UpdatedTime *time.Time     `gorm:"type:datetime null;" json:"updated_time" COMMENT:"更新时间"` // 更新时间
	DeletedTime gorm.DeletedAt `gorm:"index" json:"deleted_time" COMMENT:"删除时间"`
	Username    string         `gorm:"type:varchar(30)" json:"user_name" COMMENT:"用户账号"`
	Nickname    string         `gorm:"type:varchar(30)" json:"nick_name" COMMENT:"用户昵称"`
	Password    string         `gorm:"type:varchar(100)" json:"pass_word" COMMENT:"密码"`
	UserType    string         `gorm:"type:varchar(2)" json:"user_type" COMMENT:"用户类型（00系统用户）"`
	Email       string         `gorm:"type:varchar(50)" json:"email" COMMENT:"邮箱"`
	PhoneNumber string         `gorm:"type:varchar(11)" json:"phone_number" COMMENT:"手机"`
	Sex         string         `gorm:"type:varchar(1)" json:"sex" COMMENT:"用户性别（0男 1女 2未知）"`
	Avatar      string         `gorm:"type:varchar(100)" json:"avatar" COMMENT:"头像地址"`
	Status      bool           `gorm:"type:tinyint" json:"status" COMMENT:"帐号状态（0正常 1停用）"`
	LoginIp     string         `gorm:"type:varchar(128)" json:"login_ip" COMMENT:"最后登录IP"`
	LoginDate   *time.Time     `gorm:"type:datetime null" json:"login_date" COMMENT:"最后登录时间"`
	CreateBy    string         `gorm:"type:varchar(64)" json:"create_by" COMMENT:"创建人"`
	Remark      string         `gorm:"type:varchar(500)" json:"remark" COMMENT:"备注"`
}
