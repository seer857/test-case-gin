package system

import (
	"test-case-gin/global"
	"test-case-gin/model/System"
	"test-case-gin/utils/errmsg"
)

type MenuService struct{}

// CreateSystemMenu 新增菜单
func (menuService *MenuService) CreateSystemMenu(sysMenu *System.SysMenu) (int, *System.SysMenu) {
	err := global.GVA_DB.Create(&sysMenu).Error
	if err != nil {
		return errmsg.ERROR, sysMenu // 500
	}
	return errmsg.SUCCESS, sysMenu
}

// DeleteSystemMenu 删除菜单
func (menuService *MenuService) DeleteSystemMenu(id string) int {
	err := global.GVA_DB.Where("menu_id = ?", id).Delete(&System.SysMenu{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetAllSystemMenu 查询全部菜单
func (menuService *MenuService) GetAllSystemMenu(pageSize int, pageNum int, name string) ([]System.SysMenu, int64) {
	var documents []System.SysMenu
	var total int64

	db := global.GVA_DB

	if name != "" {
		db = db.Where("menu_name LIKE ?", "%"+name+"%")
	}

	if pageNum == 0 && pageSize == 0 {
		err := db.Find(&documents).Error
		db.Model(&System.SysMenu{}).Count(&total)
		if err != nil {
			return nil, 0
		}
		return documents, total
	} else {
		err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&documents).Error
		db.Model(&System.SysMenu{}).Count(&total)
		if err != nil {
			return nil, 0
		}
		return documents, total
	}
}

// EditSystemMenu 编辑菜单
func (menuService *MenuService) EditSystemMenu(menuId string, menu System.SysMenu) int {
	updates := map[string]interface{}{
		"menu_name": menu.MenuName,
		"parent_id": menu.ParentId,
		"order_num": menu.OrderNum,
		"path":      menu.Path,
		"component": menu.Component,
		"query":     menu.Query,
		"is_frame":  menu.IsFrame,
		"is_cache":  menu.IsCache,
		"menu_type": menu.MenuType,
		"visible":   menu.Visible,
		"status":    menu.Status,
		"perms":     menu.Perms,
		"icon":      menu.Icon,
		"remark":    menu.Remark,
	}

	// 补全代码
	err := global.GVA_DB.Model(&System.SysMenu{}).Where("menu_id = ?", menuId).Updates(updates).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
