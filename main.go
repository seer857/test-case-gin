package main

import (
	"database/sql"
	"test-case-gin/core"
	"test-case-gin/global"
	"test-case-gin/initialize"
)

func main() {
	global.GVA_DB = initialize.GormMysql()

	global.GVA_LOG = core.Zap()
	if global.GVA_DB != nil {
		initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer func(db *sql.DB) {
			err := db.Close()
			if err != nil {
				global.GVA_LOG.Info("数据库链接关闭")
			}
		}(db)
	}
	initialize.Routers()
}
