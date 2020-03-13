package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mjah/jwt-auth/api/auth"
	"github.com/spf13/viper"
)

// GetRouter sets up and handles the routes.
func GetRouter() http.Handler {
	if viper.GetString("environment") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	config := cors.DefaultConfig()

	if viper.GetBool("cors.allow_all_origins") {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = viper.GetStringSlice("cors.allow_origins")
	}

	if viper.GetBool("cors.allow_credentials") {
		config.AllowCredentials = true
	}

	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	public := r.Group("/v1")
	publicAuth := public.Group("/auth")
	{
		publicAuth.POST("/signup", auth.SignUp)
		publicAuth.POST("/signin", auth.SignIn)
		publicAuth.POST("/confirm_email", auth.ConfirmEmail)
		publicAuth.POST("/reset_password", auth.ResetPassword)
		publicAuth.POST("/send_confirm_email", auth.SendConfirmEmail)
		publicAuth.POST("/send_reset_password", auth.SendResetPasswordEmail)
	}

	private := r.Group("/v1")
	private.Use(ValidateRefreshTokenMiddleware())
	privateAuth := private.Group("/auth")
	{
		privateAuth.GET("/user_details", auth.UserDetails)
		privateAuth.GET("/signout", auth.SignOut)
		privateAuth.GET("/signout_all", auth.SignOutAll)
		privateAuth.GET("/refresh_token", auth.RefreshToken)
		privateAuth.PATCH("/update", auth.Update)
		privateAuth.DELETE("/delete", auth.Delete)
	}

	return r
}
