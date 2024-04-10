package menu

import "test-case-gin/service"

type ApiGroup struct {
	SysMenuApiApi
}

var (
	MenuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
)
