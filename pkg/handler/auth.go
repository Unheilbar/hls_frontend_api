package handler

import (
	"strings"

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

	var response int

	if detectArchiveAuth(Uri) {
		response, _ = h.services.GetResponseCodeArchive(userIp)
	} else {
		response, _ = h.services.GetResponseCodeChannel(userIp, channelAllias)
	}

	c.Status(response)
}

func detectArchiveAuth(uri string) bool {
	splitted := strings.Split(uri, "/")
	for i, val := range splitted {
		if val == "timeshift" || (val == "playlist" && i < len(splitted)-1 && splitted[i+1] == "program") {
			return true
		}
	}
	return false
}
