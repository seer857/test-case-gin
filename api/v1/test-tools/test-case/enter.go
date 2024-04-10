package test_case

import "test-case-gin/service"

type ApiGroup struct {
	TestCaseApiApi
}

var (
	testCaseService    = service.ServiceGroupApp.TestCaseServiceGroup.TestCaseService
	projectCaseService = service.ServiceGroupApp.ProjectServiceGroup.ProjectService
)
