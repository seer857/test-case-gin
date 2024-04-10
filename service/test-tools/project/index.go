package project

import (
	"test-case-gin/global"
	"test-case-gin/model/test-tools/Project"
	"test-case-gin/utils/errmsg"

	"gorm.io/gorm"
)

type ProjectService struct{}

// GetFirstProject 根据id查询名称
func (projectService *ProjectService) GetFirstProject(id string) (Project.Project, int64) {
	var project Project.Project
	err := global.GVA_DB.Where("id = ?", id).Find(&project).Error
	if err != nil {
		return project, errmsg.ERROR
	}
	return project, errmsg.SUCCESS
}

// CreateProject 新增项目
func (projectService *ProjectService) CreateProject(project *Project.Project) int {

	err := global.GVA_DB.Create(&project).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetAllProject 查询全部项目
func (projectService *ProjectService) GetAllProject(pageSize int, pageNum int) ([]Project.Project, int64) {

	var project []Project.Project
	var total int64
	err := global.GVA_DB.Find(&project).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	global.GVA_DB.Model(&project).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return project, total
}

// DeleteProject 删除项目
func (projectService *ProjectService) DeleteProject(id string) int {

	var project Project.Project
	err := global.GVA_DB.Where("id = ? ", id).Delete(&project).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// UpdateProject 编辑项目
func (projectService *ProjectService) UpdateProject(project Project.Project) int {

	var oldProject Project.Project
	if err := global.GVA_DB.First(&oldProject, "id = ?", project.ID).Error; err != nil {
		// 查询不到旧项目则返回错误码
		return errmsg.ERROR
	}
	oldProject.Name = project.Name
	oldProject.Description = project.Description
	oldProject.SerialNum = project.SerialNum
	oldProject.Category = project.Category
	// 更新到数据库
	if err := global.GVA_DB.Save(&oldProject).Error; err != nil {
		// 更新失败则返回错误码
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
