package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/unheilbar/hls_frontend_api"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
	"github.com/unheilbar/hls_frontend_api/pkg/handler"
	"github.com/unheilbar/hls_frontend_api/pkg/service"
)

func main() {
	srv := new(hls_frontend_api.Server)

	cache := cache.NewCache()

	services := service.NewService(cache)

	handlers := handler.NewHandler(services)

	go func() {
		if err := srv.Run("3000", handlers.InitRoutes()); err != nil {
			fmt.Printf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
