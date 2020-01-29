package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/api/auth"
	"github.com/spf13/viper"
)

// GetRouter ...
func GetRouter() http.Handler {
	if viper.GetString("environment") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.Default())

	apiVersion := r.Group("/v1")

	authGroup := apiVersion.Group("/auth")
	{
		authGroup.GET("/signout", auth.SignOut)
		authGroup.POST("/signin", auth.SignIn)
		authGroup.POST("/signup", auth.SignUp)
		authGroup.POST("/refreshtoken", auth.RefreshToken)
		authGroup.POST("/resetpassword", auth.ResetPassword)
		authGroup.PATCH("/confirm", auth.Confirm)
		authGroup.PATCH("/update", auth.Update)
		authGroup.DELETE("/delete", auth.Delete)
	}

	return r
}
