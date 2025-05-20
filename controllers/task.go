package controllers

import (
	"github.com/gin-gonic/gin"
	"go_to_list/db"
	"go_to_list/models"
	"net/http"
)

// 添加任务
func AddTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "add task failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "add task success", "task": task})
}
