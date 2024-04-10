package project_document

import (
	"test-case-gin/global"
	"test-case-gin/model/test-tools/ProjectDocument"
	"test-case-gin/utils/errmsg"

	"gorm.io/gorm"
)

type ProjectDocumentService struct{}

// GetAllDocument 查询全部文档目录
func (p *ProjectDocumentService) GetAllDocument(projectId string) ([]ProjectDocument.ProjectDocument, int64) {

	var cate []ProjectDocument.ProjectDocument
	var total int64
	err := global.GVA_DB.Where("project_id = ?", projectId).Order("sort").Find(&cate).Error
	global.GVA_DB.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}

// CreateProjectDocument 新增项目
func (p *ProjectDocumentService) CreateProjectDocument(project *ProjectDocument.ProjectDocument) int {

	err := global.GVA_DB.Create(&project).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// DeleteProjectDocument 删除项目
func (p *ProjectDocumentService) DeleteProjectDocument(id string) int {

	var projectDocument ProjectDocument.ProjectDocument
	err := global.GVA_DB.Where("id = ? ", id).Delete(&projectDocument).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
