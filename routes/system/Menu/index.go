package Menu

import (
	v1 "test-case-gin/api/v1"

	"github.com/gin-gonic/gin"
)

type SysMenuRouter struct{}

func (s *SysMenuRouter) InitSysMenuRouter(Router *gin.RouterGroup) {
	sysMenuRouter := Router.Group("/sys")
	sysMenuApi := v1.ApiGroupApp.SystemMenuGroup.SysMenuApiApi
	{
		sysMenuRouter.POST("/menu/add", sysMenuApi.AddMenu)
		sysMenuRouter.GET("/menu/list", sysMenuApi.GetAllMenu)
		sysMenuRouter.DELETE("/menu/del", sysMenuApi.DelMenu)
		sysMenuRouter.PUT("/menu/update", sysMenuApi.EditMenu)
	}
}
