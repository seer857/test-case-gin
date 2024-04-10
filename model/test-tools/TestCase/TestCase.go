package TestCase

import (
	"gorm.io/gorm"
	"time"
)

type TestCase struct {
	ID             int            `gorm:"int" json:"id"`
	Name           string         `gorm:"type:varchar(255)" json:"name"`
	CaseNum        string         `gorm:"type:varchar(255)" json:"caseNum"`
	CaseName       string         `gorm:"type:varchar(255)" json:"caseName"`
	CaseLevel      string         `gorm:"type:varchar(255)" json:"caseLevel"`
	CaseType       string         `gorm:"type:varchar(255)" json:"caseType"`
	Prerequisites  string         `gorm:"type:varchar(255)" json:"prerequisites"`
	TestProcedure  string         `gorm:"type:varchar(255)" json:"testProcedure"`
	ExpectedResult string         `gorm:"type:varchar(255)" json:"expectedResult"`
	Remark         string         `gorm:"type:varchar(255)" json:"remark"`
	Type           string         `gorm:"type:varchar(255)" json:"type"`
	ProjectId      string         `gorm:"type:varchar(255)" json:"projectId"`
	ProjectName    string         `gorm:"-" json:"projectName"`
	CreatedAt      time.Time      `json:"created_at"` // 创建时间
	UpdatedAt      time.Time      `json:"updated_at"` // 更新时间
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type ExportTestCase struct {
	ID             int    `gorm:"int" json:"id"`
	Name           string `gorm:"type:varchar(255)" json:"name"`
	CaseNum        string `gorm:"type:varchar(255)" json:"caseNum"`
	CaseName       string `gorm:"type:varchar(255)" json:"caseName"`
	CaseLevel      string `gorm:"type:varchar(255)" json:"caseLevel"`
	CaseType       string `gorm:"type:varchar(255)" json:"caseType"`
	Prerequisites  string `gorm:"type:varchar(255)" json:"prerequisites"`
	TestProcedure  string `gorm:"type:varchar(255)" json:"testProcedure"`
	ExpectedResult string `gorm:"type:varchar(255)" json:"expectedResult"`
	Remark         string `gorm:"type:varchar(255)" json:"remark"`
}
