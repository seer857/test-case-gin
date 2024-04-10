package login

import "test-case-gin/service"

type ApiGroup struct {
	LoginApiApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
