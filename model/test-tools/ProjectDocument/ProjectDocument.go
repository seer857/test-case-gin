package ProjectDocument

import (
	"test-case-gin/utils"
	"time"

	"gorm.io/gorm"
)

type ProjectDocument struct {
	ID        string         `gorm:"type:varchar(255)" json:"Id"`
	Name      string         `gorm:"type:varchar(255)" json:"Name"`
	Property  string         `gorm:"type:varchar(255)" json:"property"`
	ParentId  string         `gorm:"type:varchar(255)" json:"parentId"`
	ProjectId string         `gorm:"type:varchar(255)" json:"projectId"`
	Sort      int            `gorm:"type:int" json:"sort"`
	CreateBy  string         `gorm:"type:varchar(255)" json:"createBy"` // 创建人
	CreatedAt time.Time      `json:"created_at"`                        // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                        // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ProjectDocumentTree struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	Property  string                `json:"property"`
	ParentId  string                `json:"parent_id"`
	ProjectId string                `json:"projectId"`
	Children  []ProjectDocumentTree `json:"children"`
}

// BeforeCreate 钩子函数 创建之前
func (u *ProjectDocument) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = utils.RandAllString(32)
	return nil
}
