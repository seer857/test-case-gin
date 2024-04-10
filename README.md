# 测试用例生成工具

## 项目介绍

测试用例生成工具 后台代码

## 软件架构

软件架构说明

## 使用教程

1.  在根目录下创建config文件夹，并在config文件夹下创建config.ini文件

2.  在config.ini文件中配置数据库连接信息

```ini
[server]
# release 生产模式 debug 开发模式
AppMode = debug
HttpPort = :3001
[database]
Db = mysql
DbHost = 你的数据库地址
DbPort = 你的数据库端口
DbUser = 你的数据库用户名
DbPassWord = 你的数据库密码
DbName = 
[log]
MaxSize = 500
MaxBackups = 3
MaxAge = 28
Compress = false
```

3.  运行main.go文件

## 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


