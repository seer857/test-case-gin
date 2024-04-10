package files

import "test-case-gin/service"

type ApiGroup struct {
	FilesApiApi
}

var (
	filesService = service.ServiceGroupApp.FilesServiceGroup.FilesService
)
