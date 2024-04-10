package login

import (
	"net/http"
	"strconv"
	"test-case-gin/core"
	"test-case-gin/model/System"
	"test-case-gin/utils/errmsg"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginApiApi struct{}

func (login *LoginApiApi) Login(c *gin.Context) {
	var loginForm System.SysUser
	_ = c.ShouldBindJSON(&loginForm)
	var token string
	formData, code := userService.CheckLogin(loginForm.Username, loginForm.Password)
	if code == errmsg.SUCCESS {
		generateToken(c, formData)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.UserID,
			"message": errmsg.GetErrMsg(code),
			"token":   token,
		})
	}
}

// token生成函数
func generateToken(c *gin.Context, user System.SysUser) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := core.NewJWT()

	// 构造用户claims信息(负荷)
	claims := core.CustomClaims{
		Name: strconv.Itoa(user.UserID),
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "middle_school", // 签名颁发者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": errmsg.GetErrMsg(errmsg.ERROR),
			"token":   token,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"account": user.Username,
		"id":      user.UserID,
		"message": errmsg.GetErrMsg(200),
		"token":   token,
	})
	return

}
