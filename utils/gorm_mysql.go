package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

type Mysql struct {
	AppMode    string
	HttpPort   string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

var Config Mysql

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径", err)
	}
	fmt.Println("配置文件读取成功")
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	Config.AppMode = file.Section("server").Key("AppMode").MustString("debug")
	Config.HttpPort = file.Section("server").Key("HttpPort").MustString(":3001")
	Config.JwtKey = file.Section("server").Key("JwtKey").MustString("")
}
func LoadData(file *ini.File) {
	Config.Db = file.Section("database").Key("Db").MustString("mysql")
	Config.DbHost = file.Section("database").Key("DbHost").MustString("")
	Config.DbPort = file.Section("database").Key("DbPort").MustString("")
	Config.DbUser = file.Section("database").Key("DbUser").MustString("")
	Config.DbPassWord = file.Section("database").Key("DbPassWord").MustString("")
	Config.DbName = file.Section("database").Key("DbName").MustString("")
}

func (c *Mysql) Dsn() string {
	return Config.DbUser + ":" + Config.DbPassWord + "@tcp(" + Config.DbHost + ":" + Config.DbPort + ")/" + Config.DbName + "?charset=utf8&parseTime=true"
}
