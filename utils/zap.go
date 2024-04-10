package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type zapLogger struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var ZapConfig zapLogger

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	LoadLog(file)
}
func LoadLog(file *ini.File) {
	ZapConfig.MaxSize = file.Section("server").Key("MaxSize").MustInt(500)
	ZapConfig.MaxBackups = file.Section("server").Key("MaxBackups").MustInt(3)
	ZapConfig.MaxAge = file.Section("server").Key("MaxAge").MustInt(28)
	ZapConfig.Compress = file.Section("server").Key("Compress").MustBool(true)
}
