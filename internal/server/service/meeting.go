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
	in := new(MeetingCreateRequest)
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
		BeginAt:  time.UnixMilli(in.BeginAt),
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

func MeetingEdit(c *gin.Context) {
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	in := new(MeetingEditRequest)
	if err := c.ShouldBindJSON(in); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	if err := models.DB.Model(new(models.RoomBasic)).Where("identity = ? AND create_id = ?", in.Identity, uc.Id).
		Updates(map[string]any{
			"name":     in.Name,
			"begin_at": time.UnixMilli(in.BeginAt),
			"end_at":   time.UnixMilli(in.EndAt),
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

func MeetingList(c *gin.Context) {
	in := new(MeetingListRequest)
	if err := c.ShouldBindQuery(in); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}
	var list []*MeetingListReply
	var cnt int64
	tx := models.DB.Model(&models.RoomBasic{})
	if in.Keyword != "" {
		tx.Where("name LIKE ?", "%"+in.Keyword+"%")
	}
	if err := tx.Count(&cnt).Limit(in.Size).Offset((in.Page - 1) * in.Size).
		Find(&list).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "System error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

func MeetingDelete(c *gin.Context) {
	identity := c.Query("identity")
	uc := c.MustGet("user_claims").(*helper.UserClaims)
	if err := models.DB.Where("identity = ? AND create_id = ?", identity, uc.Id).
		Delete(&models.RoomBasic{}).Error; err != nil {
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
