package main

import (
	_ "github.com/devxp-tech/golang434/config"
	"github.com/devxp-tech/golang434/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()
	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  true,
			"message": "Hello world for your app golang434",
		})
	})
	server.GET("/health-check/liveness", controllers.HealthCheckLiveness)
	server.GET("/health-check/readiness", controllers.HealthCheckReadiness)
	server.Run()
}
