package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateChannelsCache(c *gin.Context) {
	err := h.services.UpdateChannelsCache()
	if err != nil {
		c.Status(500)
		return
	}
	c.Status(200)
}

func (h *Handler) clearUserCacheById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.Status(http.StatusBadRequest)
	}

	h.services.ClearUserCacheByUid(id)
	c.Status(200)
}
