package controllers

import (
	"github.com/gin-gonic/gin"
	"go_to_list/db"
	"go_to_list/models"
	"net/http"
)

// 创建清单（模拟从JWT获取UserID，实际需集成鉴权）
func CreateList(c *gin.Context) {
	var list models.List
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	list.UserID = 1 // 模拟用户ID，实际应从登录状态获取
	if err := db.DB.Create(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "create list failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "create list success", "list": list})
}

// 获取用户清单
func GetLists(c *gin.Context) {
	var lists []models.List
	if err := db.DB.Where("user_id = ?", 1).Find(&lists).Error; err != nil { // 模拟用户ID
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get lists failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lists": lists})
}
