package initialize

import (
	"os"
	"test-case-gin/global"
	"test-case-gin/model/System"
	"test-case-gin/model/test-tools/Files"
	"test-case-gin/model/test-tools/Project"
	"test-case-gin/model/test-tools/ProjectDocument"
	"test-case-gin/model/test-tools/TestCase"
)

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		//项目管理
		Project.Project{},
		ProjectDocument.ProjectDocument{},
		TestCase.TestCase{},
		Files.Files{},

		// 系统管理
		System.SysUser{},
		System.SysMenu{},
	)
	if err != nil {
		global.GVA_LOG.Info("register table success")
		os.Exit(0)
		return
	}
}
