package global

import (
	"test-case-gin/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type GvaModel struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      `json:"created_at"`           // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`           // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG utils.Server
	GVA_LOG    *zap.SugaredLogger
	GVA_TOKEN  string
)
