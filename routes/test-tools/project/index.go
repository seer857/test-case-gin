package project

import (
	v1 "test-case-gin/api/v1"

	"github.com/gin-gonic/gin"
)

type ProjectRouter struct{}

func (s *ProjectRouter) InitProjectRouter(Router *gin.RouterGroup) {
	projectRouter := Router.Group("/project")
	projectApi := v1.ApiGroupApp.ProjectApiGroup.ProjectApiApi
	projectDocumentApi := v1.ApiGroupApp.ProjectDocumentApiGroup.ProjectDocumentApiApi
	{
		// 项目管理
		projectRouter.POST("/add", projectApi.AddProject)
		projectRouter.GET("/all", projectApi.GetProject)
		projectRouter.GET("/delete", projectApi.DelProject)
		projectRouter.PUT("/update", projectApi.EditProject)
		// 项目文档目录
		projectRouter.GET("/document", projectDocumentApi.SelectDocumentTree)
		projectRouter.POST("/document/add", projectDocumentApi.AddProjectDocument)
		projectRouter.GET("/document/delete", projectDocumentApi.DelProjectDocument)
	}
}
