package user

import (
	"net/http"
	"strconv"
	"test-case-gin/model/System"
	"test-case-gin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type SysUserApiApi struct{}

// AddUser 添加用户
func (sysUserApiApi *SysUserApiApi) AddUser(c *gin.Context) {
	var data System.SysUser
	_ = c.ShouldBindJSON(&data)

	code, User := userService.CreateSystemUser(&data)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    User,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

type listQuery struct {
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	Status      string `json:"status"`
	CreateTime  string `json:"createTime"`
	PageSize    int    `json:"pageSize"`
	PageNum     int    `json:"pageNum"`
}

// GetUserList 获取用户列表
func (sysUserApiApi *SysUserApiApi) GetUserList(c *gin.Context) {

	var params listQuery
	params.Username = c.Query("username")
	params.PhoneNumber = c.Query("phoneNumber")
	params.Status = c.Query("status")
	params.CreateTime = c.Query("createTime")
	params.PageSize, _ = strconv.Atoi(c.Query("pageSize"))
	params.PageNum, _ = strconv.Atoi(c.Query("pageNum"))

	data, total, code := userService.SelectLikeUser(
		params.PageSize,
		params.PageNum,
		params.Username,
		params.PhoneNumber,
		params.Status,
	)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
			"params":  params,
		},
	)
}
