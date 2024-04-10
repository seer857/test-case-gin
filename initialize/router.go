package initialize

import (
	v1 "test-case-gin/api/v1"
	"test-case-gin/core"
	"test-case-gin/global"
	"test-case-gin/routes"
	"test-case-gin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(gin.Recovery())
	Router.Use(core.Cors())
	PrivateGroup := Router.Group("/speediness")
	// 静态资源
	PrivateGroup.Static("/uploads", "./uploads")
	// 登录
	loginApi := v1.ApiGroupApp.LoginApiGroup.LoginApiApi
	PrivateGroup.POST("/login", loginApi.Login)
	// 文件上传
	fileApi := v1.ApiGroupApp.FilesApiGroup.FilesApiApi
	PrivateGroup.POST("/upload/excel", fileApi.UploadFile)

	// 测试用例
	TestCaseRouter := routes.RouterGroupApp.TestCase
	TestCaseRouter.InitTestCaseRouter(PrivateGroup)
	// 项目管理
	ProjectRouter := routes.RouterGroupApp.Project
	ProjectRouter.InitProjectRouter(PrivateGroup)
	// 文件管理
	FileRouter := routes.RouterGroupApp.Files
	FileRouter.InitFilesRouter(PrivateGroup)
	// 系统管理 - 用户管理
	SysUserRouter := routes.RouterGroupApp.SysUser
	SysUserRouter.InitSysUserRouter(PrivateGroup)
	SysMenuRouter := routes.RouterGroupApp.SysMenu
	SysMenuRouter.InitSysMenuRouter(PrivateGroup)

	global.GVA_LOG.Info("server run success on ", zap.String("address", utils.Config.HttpPort))
	_ = Router.Run(utils.Config.HttpPort)
	return Router
}
