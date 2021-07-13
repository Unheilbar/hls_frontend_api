package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateChannelsCache(c *gin.Context) {
	h.services.UpdateChannelsCache()
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
