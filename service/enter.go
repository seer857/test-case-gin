package service

import (
	"test-case-gin/service/system"
	"test-case-gin/service/test-tools/files"
	"test-case-gin/service/test-tools/project"
	projectdocument "test-case-gin/service/test-tools/project-document"
	testcase "test-case-gin/service/test-tools/test-case"
)

type GroupService struct {
	TestCaseServiceGroup        testcase.TestCaseServiceGroup
	ProjectServiceGroup         project.ProjectServiceGroup
	ProjectDocumentServiceGroup projectdocument.ProjectDocumentServiceGroup
	FilesServiceGroup           files.FilesServiceGroup
	SystemServiceGroup          system.SystemServiceGroup
}

var ServiceGroupApp = new(GroupService)
