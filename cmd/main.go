package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
	"github.com/unheilbar/hls_frontend_api/pkg/handler"
	"github.com/unheilbar/hls_frontend_api/pkg/service"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))

	mode := os.Getenv("mode")
	if mode == "release" {
		logrus.SetLevel(logrus.InfoLevel)
	}

	srv := new(hls_frontend_api.Server)

	cache := cache.NewCache()

	services := service.NewService(cache)

	handlers := handler.NewHandler(services)

	go func() {
		if err := srv.Run(os.Getenv("port"), handlers.InitRoutes()); err != nil {
			logrus.Printf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
