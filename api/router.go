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

	if viper.GetBool("serve.cors.allow_all_origins") {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = viper.GetStringSlice("serve.cors.allow_origins")
	}

	if viper.GetBool("serve.cors.allow_credentials") {
		config.AllowCredentials = true
	}

	config.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	public := r.Group("/")
	{
		public.POST("signup", auth.SignUp)
		public.POST("signin", auth.SignIn)
		public.POST("confirm-email", auth.ConfirmEmail)
		public.POST("reset-password", auth.ResetPassword)
		public.POST("send-confirm-email", auth.SendConfirmEmail)
		public.POST("send-reset-password", auth.SendResetPasswordEmail)
	}

	private := r.Group("/")
	private.Use(ValidateRefreshTokenMiddleware())
	{
		private.GET("user", auth.UserDetails)
		private.PATCH("user", auth.Update)
		private.DELETE("user", auth.Delete)
		private.GET("signout", auth.SignOut)
		private.GET("signout-all", auth.SignOutAll)
		private.GET("refresh-token", auth.RefreshToken)
	}

	return r
}
