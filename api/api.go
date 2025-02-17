package api

import (
	"chx-passport/api/middleware"
	"chx-passport/config"
	"chx-passport/controller"
	"log"

	"github.com/gin-gonic/gin"
)

func RunApi() {
	gin.SetMode(config.ConfigContext.ApiConfig.Mode)
	host := config.ConfigContext.ApiConfig.Host
	port := config.ConfigContext.ApiConfig.Port
	log.Println("Starting API on http://" + host + ":" + port)
	r := gin.Default()

	r.Use(middleware.Cors()).Use(middleware.ShowUserInfo()) // 允许跨域

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Caillo World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	r.POST("/refresh", controller.RefreshToken)

	needLoginGroup := r.Group("/user", middleware.Auth())
	{
		needLoginGroup.GET("/info", controller.UserInfo)
		needLoginGroup.POST("/change_info", controller.ChangeInfo)
		// needLoginGroup.POST("/update", controller.UpdateUserInfo)
	}

	r.Run(host + ":" + port)
}
