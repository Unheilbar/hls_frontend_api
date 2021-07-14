package handler

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

const (
	aliasHeader     = "X-Original-URI"
	realIpHeader    = "X-Real-IP"
	forwardedHeader = "X-Forwarded-For"
	UriHeader       = "Requested-Uri"
)

func (h *Handler) auth(c *gin.Context) {
	channelAllias := c.GetHeader(aliasHeader)

	if channelAllias == "" {
		channelAllias = "hz"
	}

	userIp := c.GetHeader(realIpHeader)

	if userIp == "" {
		userIp = c.GetHeader(forwardedHeader)
	}

	Uri := c.GetHeader(UriHeader)

	r, _ := regexp.Compile(`/playlist/program/.*m3u8`)

	var response int
	if r.MatchString(Uri) {
		response, _ = h.services.GetResponseCodeArchive(userIp)
	} else {
		response, _ = h.services.GetResponseCodeChannel(userIp, channelAllias)
	}

	c.Status(response)
}
