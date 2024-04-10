package files

import (
	"test-case-gin/global"
	"test-case-gin/model/test-tools/Files"
	"test-case-gin/utils/errmsg"

	"gorm.io/gorm"
)

type FilesService struct{}

// CreateFiles 创建文件
func (filesService *FilesService) CreateFiles(file *Files.Files) int {
	err := global.GVA_DB.Create(&file).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetAllFiles 查询全部文件
func (filesService *FilesService) GetAllFiles(pageSize int, pageNum int) ([]Files.Files, int64) {
	var files []Files.Files
	var total int64
	err := global.GVA_DB.Find(&files).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	global.GVA_DB.Model(&files).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return files, total
}

// GetUploadTypeFiles 按类型查询文件
func (filesService *FilesService) GetUploadTypeFiles(pageSize int, pageNum int, uploadType string) ([]Files.Files, int64) {
	var files []Files.Files
	var total int64
	global.GVA_DB.Model(&Files.Files{})
	err := global.GVA_DB.Where("upload_type = ?", uploadType).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&files).Error
	global.GVA_DB.Count(&total)
	if err != nil {
		return nil, 0
	}
	return files, total
}

// SelectFile 查询单个文件
func (filesService *FilesService) SelectFile(id string) (Files.Files, int) {
	var files Files.Files
	err := global.GVA_DB.Where("id = ?", id).First(&files).Error
	if err != nil {
		return files, 500
	}
	return files, 200
}

// DeleteFile 删除文件
func (filesService *FilesService) DeleteFile(id int) (Files.Files, int) {
	var files Files.Files
	// 查询文件信息
	global.GVA_DB.Where("id = ?", id).Find(&files)
	err := global.GVA_DB.Where("id = ?", id).Delete(&files).Error
	if err != nil {
		return files, errmsg.ERROR
	}

	return files, errmsg.SUCCESS
}

// UpdateFiles 更新文件状态
func (filesService *FilesService) UpdateFiles(files *Files.Files) int {
	var oldFiles Files.Files
	if err := global.GVA_DB.First(&oldFiles, "id = ?", files.ID).Error; err != nil {
		// 查询不到旧项目则返回错误码
		return errmsg.ERROR
	}
	oldFiles.FileNameZh = files.FileNameZh
	oldFiles.FileStatus = files.FileStatus
	// 更新到数据库
	if err := global.GVA_DB.Save(&oldFiles).Error; err != nil {
		// 更新失败则返回错误码
		return errmsg.ERROR
	}

	return errmsg.SUCCESS
}
