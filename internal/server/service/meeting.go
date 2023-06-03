package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"vediomeeting/internal/helper"
	"vediomeeting/internal/models"
)

func MeetingCreate(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	in := new(MeetCreateRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if err := models.DB.Create(&models.RoomBasic{
		Identity: helper.GetUUID(),
		Name:     in.Name,
		BeginAt:  time.UnixMilli(in.CreateAt),
		EndAt:    time.UnixMilli(in.EndAt),
		CreateId: uc.Id,
	}).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "System error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
	})
}
