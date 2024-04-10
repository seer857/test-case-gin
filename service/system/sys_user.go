package system

import (
	"test-case-gin/global"
	"test-case-gin/model/System"
	"test-case-gin/utils/errmsg"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// CheckLogin 检查用户是否存在
func (userService *UserService) CheckLogin(username string, password string) (System.SysUser, int) {
	var user System.SysUser
	var PasswordErr error

	global.GVA_DB.Where("username = ?", username).First(&user)

	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.UserID == -1 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCESS
}

// CreateSystemUser 新增用户
func (userService *UserService) CreateSystemUser(sysUser *System.SysUser) (int, *System.SysUser) {

	err := global.GVA_DB.Create(&sysUser).Error
	if err != nil {
		return errmsg.ERROR, sysUser // 500
	}
	return errmsg.SUCCESS, sysUser
}

// SelectLikeUser 根据关键词模糊查询接口
func (userService *UserService) SelectLikeUser(pageSize int, pageNum int, username string, phoneNumber string, status string) ([]System.SysUser, int64, int) {

	var sysUser []System.SysUser
	var total int64
	db := global.GVA_DB.Model(&System.SysUser{})
	if username != "" {
		db = db.Where("username LIKE ?", "%"+username+"%")
	}
	if phoneNumber != "" {
		db = db.Where("phone_number LIKE ?", "%"+phoneNumber+"%")
	}
	if status != "" {
		db = db.Where("status LIKE ?", "%"+status+"%")
	}
	err := db.Order("user_id").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&sysUser).Error
	db.Count(&total)
	if err != nil {
		return sysUser, 0, errmsg.ERROR
	}
	return sysUser, total, errmsg.SUCCESS
}
