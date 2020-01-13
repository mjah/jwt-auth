package auth

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Run ...
func Run() {
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/v1")

	auth := api.Group("/auth")
	{
		auth.GET("/logout", Logout)
		auth.POST("/login", Login)
		auth.POST("/signup", SignUp)
		auth.POST("/resetpassword", ResetPassword)
		auth.PUT("/confirm", Confirm)
		auth.PUT("/update", Update)
		auth.DELETE("/delete", Delete)
	}

	r.Run(":9096")
}
