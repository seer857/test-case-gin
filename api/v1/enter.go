package v1

import (

	"test-case-gin/api/v1/system/menu"
	"test-case-gin/api/v1/system/user"
	"test-case-gin/api/v1/test-tools/files"
	"test-case-gin/api/v1/test-tools/login"
	"test-case-gin/api/v1/test-tools/project"
	"test-case-gin/api/v1/test-tools/projectDocument"
	testcase "test-case-gin/api/v1/test-tools/test-case"
)

type ApiGroup struct {
	// PC端
	LoginApiGroup           login.ApiGroup
	ProjectApiGroup         project.ApiGroup
	TestCaseApiGroup        testcase.ApiGroup
	ProjectDocumentApiGroup projectDocument.ApiGroup
	FilesApiGroup           files.ApiGroup
	// 系统管理
	SystemUserGroup user.ApiGroup
	SystemMenuGroup menu.ApiGroup

}

var ApiGroupApp = new(ApiGroup)
