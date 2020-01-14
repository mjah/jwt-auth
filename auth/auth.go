package auth

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Run ...
func Run() {
	keyPath := viper.GetString("token.private_key_path")
	loadPrivateKey(&keyPath)

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/v1")

	auth := api.Group("/auth")
	{
		auth.GET("/logout", Logout)
		auth.POST("/login", Login)
		auth.POST("/signup", SignUp)
		auth.POST("/refreshtoken", RefreshToken)
		auth.POST("/resetpassword", ResetPassword)
		auth.PATCH("/confirm", Confirm)
		auth.PATCH("/update", Update)
		auth.DELETE("/delete", Delete)
	}

	r.Run(":9096")
}
