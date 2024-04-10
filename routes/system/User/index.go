package User

import (
	v1 "test-case-gin/api/v1"

	"github.com/gin-gonic/gin"
)

type SysUserRouter struct{}

func (s *SysUserRouter) InitSysUserRouter(Router *gin.RouterGroup) {
	sysUserRouter := Router.Group("/sys")
	sysUserApi := v1.ApiGroupApp.SystemUserGroup.SysUserApiApi
	{

		sysUserRouter.POST("/user/add", sysUserApi.AddUser)
		sysUserRouter.GET("/user/list", sysUserApi.GetUserList)
	}
}
