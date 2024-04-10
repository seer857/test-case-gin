package test_case

import (
	v1 "test-case-gin/api/v1"

	"github.com/gin-gonic/gin"
)

type TestCaseRouter struct{}

func (s *TestCaseRouter) InitTestCaseRouter(Router *gin.RouterGroup) {
	testCaseRouter := Router.Group("/test-case")
	testCaseApi := v1.ApiGroupApp.TestCaseApiGroup.TestCaseApiApi
	{
		testCaseRouter.GET("/add", testCaseApi.AddTestCase)
		testCaseRouter.GET("/all", testCaseApi.GetTestCase)
		testCaseRouter.DELETE("/delete", testCaseApi.DelTestCase)
		testCaseRouter.PUT("/update", testCaseApi.UpdateTestCase)
		testCaseRouter.POST("/search", testCaseApi.LikeSelectTestCase)
		// 导出测试用例 Excel
		testCaseRouter.GET("/export", testCaseApi.ExportExcel)

		// 测试用例渲染到模板
		testCaseRouter.GET("/read/word", testCaseApi.ReadWordTemplate)
		testCaseRouter.GET("/file", testCaseApi.FileBlob)
	}
}
