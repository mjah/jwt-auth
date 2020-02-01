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

	public := r.Group("/v1")
	publicAuth := public.Group("/auth")
	{
		publicAuth.POST("/signin", auth.SignIn)
		publicAuth.POST("/signup", auth.SignUp)
		publicAuth.POST("/refreshtoken", auth.RefreshToken)
		publicAuth.POST("/resetpassword", auth.ResetPassword)
		publicAuth.PATCH("/confirm", auth.Confirm)
	}

	private := r.Group("/v1")
	private.Use(ValidateRefreshTokenMiddleware())
	privateAuth := private.Group("/auth")
	{
		privateAuth.GET("/signout", auth.SignOut)
		privateAuth.PATCH("/update", auth.Update)
		privateAuth.DELETE("/delete", auth.Delete)
	}

	return r
}
