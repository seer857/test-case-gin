package routes

import (
	"test-case-gin/routes/system/Menu"
	"test-case-gin/routes/system/User"
	"test-case-gin/routes/test-tools/Files"
	"test-case-gin/routes/test-tools/project"
	test_case "test-case-gin/routes/test-tools/test-case"
)

type RouterGroup struct {
	TestCase test_case.RouterGroup
	Project  project.RouterGroup
	Files    Files.RouterGroup

	SysUser User.RouterGroup
	SysMenu Menu.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
