package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"vediomeeting/internal/helper"
	"vediomeeting/internal/models"
)

func UserLogin(c *gin.Context) {
	in := new(UserLoginRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常！",
		})
		return
	}

	if in.Username == "" || in.Password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户名或密码为空！",
		})
		return
	}

	in.Password = helper.GetMd5(in.Password)

	data := new(models.UserBasic)
	if err := models.DB.Where("username= ? AND password = ?", in.Username, in.Password).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户名或密码错误！",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Get UserBasic Error:" + err.Error(),
		})
		return
	}
	token, err := helper.GenerateToken(data.ID, data.Username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": map[string]interface{}{
			"token": token,
		},
	})
}
