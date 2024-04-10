package initialize

import (
	"fmt"
	"test-case-gin/global"
	"test-case-gin/utils"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func GormMysql() *gorm.DB {
	m := utils.Config

	if m.DbName == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `User`
			SingularTable: true,
		},
	})
	if err != nil {
		global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		fmt.Println("连接数据库失败", err)
	} else {
		fmt.Println("数据库连接成功！！")
	}
	return db
}
