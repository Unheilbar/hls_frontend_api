package handler

import (
	"os"
	"strconv"

	limit "github.com/aviddiviner/gin-limit"
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

	limit_connections, err := strconv.Atoi(os.Getenv("limit_gourutines"))

	if err != nil {
		limit_connections = 2000
	}

	router.Use(limit.MaxAllowed(limit_connections))

	router.GET("/auth", h.auth)

	router.GET("/manage/reload_cha", h.updateChannelsCache)

	router.GET("/manage/delete_user/:uid", h.clearUserCacheById)

	return router
}
