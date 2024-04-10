package projectDocument

import "test-case-gin/service"

type ApiGroup struct{ ProjectDocumentApiApi }

var (
	projectDocumentService = service.ServiceGroupApp.ProjectDocumentServiceGroup.ProjectDocumentService
)
