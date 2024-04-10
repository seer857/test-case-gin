package project

import (
	"net/http"
	"strconv"
	"test-case-gin/model/test-tools/Project"
	"test-case-gin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type ProjectApiApi struct{}

// AddProject 新增项目
func (projectApiApi *ProjectApiApi) AddProject(c *gin.Context) {
	var params Project.Project
	_ = c.ShouldBindJSON(&params)
	code := projectService.CreateProject(&params)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetProject 查询全部项目
func (projectApiApi *ProjectApiApi) GetProject(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	data, total := projectService.GetAllProject(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// DelProject 删除项目
func (projectApiApi *ProjectApiApi) DelProject(c *gin.Context) {
	id := c.Query("id")

	code := projectService.DeleteProject(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditProject 更新项目
func (projectApiApi *ProjectApiApi) EditProject(c *gin.Context) {
	var params Project.Project
	_ = c.ShouldBindJSON(&params)
	code := projectService.UpdateProject(params)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
