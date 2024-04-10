package project

import "test-case-gin/service"

type ApiGroup struct {
	ProjectApiApi
}

var (
	projectService = service.ServiceGroupApp.ProjectServiceGroup.ProjectService
)
