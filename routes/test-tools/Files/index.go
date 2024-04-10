package Files

import (
	v1 "test-case-gin/api/v1"

	"github.com/gin-gonic/gin"
)

type FilesRouter struct{}

func (s *FilesRouter) InitFilesRouter(Router *gin.RouterGroup) {
	filesRouter := Router.Group("/files")
	filesApi := v1.ApiGroupApp.FilesApiGroup.FilesApiApi
	{

		// 上传文件管理
		filesRouter.GET("/all", filesApi.GetFiles)
		filesRouter.POST("/delete", filesApi.DelFile)
		filesRouter.GET("/uploadType", filesApi.GetTypeFiles)

		// 机器学习-调用模型
		filesRouter.POST("/mode/dataInput", filesApi.DataInput)
		filesRouter.POST("/mode/details", filesApi.SelectInputData)
	}
}
