package internal

import (
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/syzhang42/hermes/utils/ormx"
)

func signIn() gin.HandlerFunc {

	return func(c *gin.Context) {
		var userinfo UserInfo
		if err := c.ShouldBindJSON(&userinfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//注册码
		if userinfo.AuthKey == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "AuthKey must have a value", "code": -1})
			return
		}
		//注册码审核
		var req = []AuthCode{{AuthKey: userinfo.AuthKey}}
		if err := ormx.GetPostgresCli().Find("authkey = ?", userinfo.AuthKey).Find(&req).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": -1})
			return
		}
		if len(req) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "注册码无效", "code": -1})
			return
		}
		if userinfo.UserName == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "UserName must have a value", "code": -1})
			return
		}
		if userinfo.Password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password must have a value", "code": -1})
			return
		}
		var res []UserInfo
		if err := ormx.GetPostgresCli().Find("username = ?", userinfo.UserName).Find(&res).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": -1})
			return
		}
		if len(res) != 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "该用户已注册", "code": -1})
			return
		}
		// 正则表达式来校验密码，必须包含大小写字母和数字，且只能是这些字符
		match := isValidPassword(userinfo.Password)
		if !match {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Password must contain at least one uppercase letter, one lowercase letter, and one digit, and must be at least 8 characters long", "code": -1})
			return
		}
		// 写入
		if err := ormx.GetPostgresCli().Create(&userinfo).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": -1})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "success", "code": 0})
	}
}

func logIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userinfo UserInfo
		if err := c.ShouldBindJSON(&userinfo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if userinfo.UserName == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "UserName must have a value", "code": -1})
			return
		}
		if userinfo.Password == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "password must have a value", "code": -1})
			return
		}
		var res []UserInfo
		if err := ormx.GetPostgresCli().Find("username = ? AND password = ?", userinfo.UserName, userinfo.Password).Find(&res).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "code": -1})
			return
		}
		if len(res) == 0 {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "密码错误", "code": -1})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "success", "code": 0, "meta": res[0].Mata})
	}
}
func isValidPassword(password string) bool {
	// 长度至少为 8
	if len(password) < 8 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasDigit := false

	// 检查密码中的字符类型
	for _, ch := range password {
		if unicode.IsLower(ch) {
			hasLower = true
		} else if unicode.IsUpper(ch) {
			hasUpper = true
		} else if unicode.IsDigit(ch) {
			hasDigit = true
		}
	}

	// 检查是否同时包含小写字母、大写字母和数字
	return hasLower && hasUpper && hasDigit
}
