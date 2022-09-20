package main

import (
	"http_monitor/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	r.GET("/register")
	r.Run()
}
