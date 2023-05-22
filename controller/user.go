package controller

import (
	"net/http"

	"afthab/github.com/config"
	"afthab/github.com/models"

	"github.com/gin-gonic/gin"
)

var Err error

func ViewAllTask(c *gin.Context) {
	new := []models.Task{}
	Err = config.DB.Find(&new).Error
	if Err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Error in getting the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"All Tasks": new,
	})
}
func AddTask(c *gin.Context) {
	new := models.Task{}
	Err = c.Bind(&new)
	if Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": Err,
		})
		return
	}
	Err = config.DB.Create(&new).Error
	if Err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": Err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": "Added",
	})

}
func UpdateTask(c *gin.Context) {
	new := models.UpdateTask{}
	Err = c.Bind(&new)
	if Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": Err,
		})
		return
	}
	Err = config.DB.Model(&models.Task{}).Where("taskname = ?", new.Oldtask).
		Updates(models.Task{Taskname: new.Newtask, Due: new.Newdue}).Error
	if Err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": Err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Updated",
	})
}

func DeleteTask(c *gin.Context) {
	new := models.Task{}
	Err = c.Bind(&new)
	if Err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": Err,
		})
		return
	}
	Err = config.DB.Delete(&models.Task{}, "taskname = ?", new.Taskname).Error
	if Err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": Err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Deleted",
	})

}
