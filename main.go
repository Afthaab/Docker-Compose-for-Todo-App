package main

import (
	"afthab/github.com/config"
	"afthab/github.com/initializers"
	"os"

	"afthab/github.com/controller"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()

func init() {
	initializers.LoadEnv()
	config.DBconnect()
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	R.GET("/task/view", controller.ViewAllTask)
	R.POST("task/add", controller.AddTask)
	R.PUT("task/edit", controller.UpdateTask)
	R.DELETE("task/delete", controller.DeleteTask)
	R.Run(os.Getenv("PORT"))
}
