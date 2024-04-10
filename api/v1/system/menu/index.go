package menu

import (
	"net/http"
	"strconv"
	"test-case-gin/model/System"
	"test-case-gin/utils/errmsg"

	"github.com/gin-gonic/gin"
)

type SysMenuApiApi struct{}

// AddMenu 新增菜单
func (p *SysMenuApiApi) AddMenu(c *gin.Context) {
	var params System.SysMenu
	_ = c.ShouldBindJSON(&params)
	code, menu := MenuService.CreateSystemMenu(&params)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    menu,
	})
}

type Params struct {
	Id string `json:"id"`
}

// DelMenu 删除菜单
func (p *SysMenuApiApi) DelMenu(c *gin.Context) {
	var params Params
	_ = c.ShouldBindJSON(&params)
	code := MenuService.DeleteSystemMenu(params.Id)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// EditMenu 编辑菜单
func (p *SysMenuApiApi) EditMenu(c *gin.Context) {
	var params System.SysMenu
	_ = c.ShouldBindJSON(&params)
	code := MenuService.EditSystemMenu(params.MenuId, params)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetAllMenu 获取所有菜单
func (p *SysMenuApiApi) GetAllMenu(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	pageNum, _ := strconv.Atoi(c.Query("page_num"))
	name := c.Query("name")
	menus, total := MenuService.GetAllSystemMenu(pageSize, pageNum, name)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  errmsg.SUCCESS,
			"message": errmsg.GetErrMsg(errmsg.SUCCESS),
			"data":    menus,
			"total":   total,
		},
	)
}
