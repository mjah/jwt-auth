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

	environment := viper.GetString("environment")

	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/v1")

	auth := api.Group("/auth")
	{
		auth.GET("/signout", SignOut)
		auth.POST("/signin", SignIn)
		auth.POST("/signup", SignUp)
		auth.POST("/refreshtoken", RefreshToken)
		auth.POST("/resetpassword", ResetPassword)
		auth.PATCH("/confirm", Confirm)
		auth.PATCH("/update", Update)
		auth.DELETE("/delete", Delete)
	}

	r.Run(":9096")
}
