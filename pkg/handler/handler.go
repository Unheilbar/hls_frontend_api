package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/unheilbar/hls_frontend_api/pkg/service"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	mode := os.Getenv("mode")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	router.GET("/auth", h.auth)

	router.GET("/streaming/reload_cha", h.updateChannelsCache)

	router.GET("/delete_user/:uid", h.clearUserCacheById)

	return router
}
