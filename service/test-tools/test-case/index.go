package test_case

import (
	"errors"
	"test-case-gin/global"
	"test-case-gin/model/test-tools/TestCase"
	"test-case-gin/utils/errmsg"

	"gorm.io/gorm"
)

type TestCaseService struct{}

// CreateTestCase 新增测试实例
func (testCaseService *TestCaseService) CreateTestCase(testCase *TestCase.TestCase) int {

	err := global.GVA_DB.Create(&testCase).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetAllTestCaseLimit 查询全部测试用例
func (testCaseService *TestCaseService) GetAllTestCaseLimit(id string, property string, pageSize int, pageNum int) ([]TestCase.TestCase, int64) {

	var testCase []TestCase.TestCase
	var total int64
	// 先计算总数
	global.GVA_DB.Model(&TestCase.TestCase{}).Count(&total)

	// 进行查询
	if pageNum == 0 {
		// 如果pageNum为0，则返回所有结果
		if property != "" {
			err := global.GVA_DB.Where("project_id = ? AND type = ?", id, property).Find(&testCase).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, 0
			}
			// 根据查询条件计算总数
			global.GVA_DB.Model(&TestCase.TestCase{}).Where("project_id = ? AND type = ?", id, property).Count(&total)
		} else {
			err := global.GVA_DB.Where("project_id = ?", id).Find(&testCase).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, 0
			}
			// 根据查询条件计算总数
			global.GVA_DB.Model(&TestCase.TestCase{}).Where("project_id = ?", id).Count(&total)
		}
	} else {
		// 否则，进行分页查询
		if property != "" {
			err := global.GVA_DB.Where("project_id = ? AND type = ?", id, property).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&testCase).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, 0
			}
			// 根据查询条件计算总数
			global.GVA_DB.Model(&TestCase.TestCase{}).Where("project_id = ? AND type = ?", id, property).Count(&total)
		} else {
			err := global.GVA_DB.Where("project_id = ?", id).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&testCase).Error
			if err != nil && err != gorm.ErrRecordNotFound {
				return nil, 0
			}
			// 根据查询条件计算总数
			global.GVA_DB.Model(&TestCase.TestCase{}).Where("project_id = ?", id).Count(&total)
		}
	}
	return testCase, total
}

func (testCaseService *TestCaseService) GetAllTestCase(pageSize int, pageNum int) ([]TestCase.TestCase, int64) {

	var testCase []TestCase.TestCase
	var total int64
	// 先计算总数
	global.GVA_DB.Model(&TestCase.TestCase{}).Count(&total)

	// 进行查询
	if pageNum == 0 {
		// 如果pageNum为0，则返回所有结果
		err := global.GVA_DB.Find(&testCase).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	} else {
		// 否则，进行分页查询
		err := global.GVA_DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&testCase).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	}
	return testCase, total
}

// DeleteTestCase 删除测试
func (testCaseService *TestCaseService) DeleteTestCase(id int) int {

	var testCase TestCase.TestCase
	err := global.GVA_DB.Where("id = ? ", id).Delete(&testCase).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// SelectLikeTestCaseId 根据关键词模糊查询接口
func (testCaseService *TestCaseService) SelectLikeTestCaseId(id string) ([]TestCase.TestCase, int64, int) {

	var testCase []TestCase.TestCase
	var total int64
	err := global.GVA_DB.Model(&testCase).Where("project_id = ? ", id).Find(&testCase).Error
	global.GVA_DB.Model(&testCase).Where("project_id = ? ", id).Find(&testCase).Count(&total)
	if err != nil {
		return testCase, 0, errmsg.ERROR
	}
	return testCase, total, errmsg.SUCCESS
}

// SelectLikeTestCaseIdProperty 根据关键词模糊查询接口
func (testCaseService *TestCaseService) SelectLikeTestCaseIdProperty(id string, property string) ([]TestCase.TestCase, int64, int) {

	var testCase []TestCase.TestCase
	var total int64
	err := global.GVA_DB.Model(&testCase).
		Where("project_id = ? AND type = ?", id, property).
		Find(&testCase).Error
	global.GVA_DB.Model(&testCase).
		Where("project_id = ? AND type = ?", id, property).
		Count(&total)
	if err != nil {
		return testCase, 0, errmsg.ERROR
	}
	return testCase, total, errmsg.SUCCESS
}

// SelectCaseTest 查询单个测试用例
func (testCaseService *TestCaseService) SelectCaseTest(id string) TestCase.TestCase {

	var testCase TestCase.TestCase
	err := global.GVA_DB.
		Where("id = ?", id).First(&testCase).Error
	if err != nil {
		return testCase
	}
	return testCase
}

// UpdateTestCase 编辑测试用例
func (testCaseService *TestCaseService) UpdateTestCase(testCase TestCase.TestCase) int {

	var oldTestCase TestCase.TestCase
	if err := global.GVA_DB.First(&oldTestCase, "id = ?", testCase.ID).Error; err != nil {
		// 查询不到旧项目则返回错误码
		return errmsg.ERROR
	}
	oldTestCase.Name = testCase.Name
	oldTestCase.CaseNum = testCase.CaseNum
	oldTestCase.CaseName = testCase.CaseName
	oldTestCase.CaseLevel = testCase.CaseLevel
	oldTestCase.CaseType = testCase.CaseType
	oldTestCase.Prerequisites = testCase.Prerequisites
	oldTestCase.TestProcedure = testCase.TestProcedure
	oldTestCase.ExpectedResult = testCase.ExpectedResult
	oldTestCase.Remark = testCase.Remark
	// 更新到数据库
	if err := global.GVA_DB.Save(&oldTestCase).Error; err != nil {
		// 更新失败则返回错误码
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
