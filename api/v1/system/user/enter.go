package user

import "test-case-gin/service"

type ApiGroup struct {
	SysUserApiApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
