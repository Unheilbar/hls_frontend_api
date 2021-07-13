package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/unheilbar/hls_frontend_api/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/auth", h.auth)

	router.GET("/streaming/reload_cha", h.updateChannelsCache)

	router.GET("/delete_user/:uid", h.clearUserCacheById)

	return router
}
