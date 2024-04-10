package Files

import (
	"gorm.io/gorm"
	"time"
)

type Files struct {
	ID         int            `gorm:"type:int" json:"id"`                  // 项目ID
	FileNameZh string         `gorm:"type:varchar(255)" json:"fileNameZH"` //文件名称
	FileName   string         `gorm:"type:varchar(255)" json:"fileName"`   //文件名称
	FileUrl    string         `gorm:"type:varchar(255)" json:"fileUrl"`    //文件存储路径
	FileType   string         `gorm:"type:varchar(255)" json:"fileType"`   // 文件类型
	FileStatus string         `gorm:"type:varchar(2)" json:"fileStatus"`   // 文件状态
	UploadType string         `gorm:"type:varchar(255)" json:"uploadType"` // 上传类型
	ProjectId  string         `gorm:"type:varchar(255)" json:"projectId"`
	CreateBy   string         `gorm:"type:varchar(255)" json:"createBy"` // 创建人
	CreatedAt  time.Time      `json:"created_at"`                        // 创建时间
	UpdatedAt  time.Time      `json:"updated_at"`                        // 更新时间
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
