package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	aliasHeader     = "x-original-uri"
	realIpHeader    = "X-Real-IP"
	forwardedHeader = "X-Forwarded-For"
)

func (h *Handler) Auth(c *gin.Context) {
	channelAllias := c.GetHeader(aliasHeader)

	if channelAllias == "" {
		channelAllias = "hz"
	}

	userIp := c.GetHeader(realIpHeader)

	if userIp == "" {
		userIp = c.GetHeader(forwardedHeader)
	}

	response, err := h.services.GetResponseCode(userIp, channelAllias)

	h.services.AddUserCacheItem(userIp, channelAllias)

	if err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.Status(response)
}

func (h *Handler) ReloadChannels(c *gin.Context) {
	h.services.ReloadChannels()
	c.Status(200)

}
