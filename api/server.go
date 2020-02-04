package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mjah/jwt-auth/logger"
	"github.com/spf13/viper"
)

// Serve ...
func Serve() {
	serverAddress := viper.GetString("serve.host") + ":" + viper.GetString("serve.port")
	srv := &http.Server{
		Addr:    serverAddress,
		Handler: GetRouter(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log().Fatal("Serve error: ", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	logger.Log().Info("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log().Fatal("Server shutdown: ", err)
	}

	logger.Log().Info("Server exiting.")
}
