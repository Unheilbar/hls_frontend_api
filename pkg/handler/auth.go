package handler

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	authQueryType, err := detectAuthQueryType(Uri)
	if err != nil {
		response = 403
	}

	// on timeshift we check field user.cacheItem.Arh and if channel code is availabe

	if authQueryType == "timeshift" {
		response, _ = h.services.GetResponseCodeChannel(userIp, channelAllias, true)
	}

	// on streaming we check only if channel code is available
	if authQueryType == "streaming" {
		response, _ = h.services.GetResponseCodeChannel(userIp, channelAllias, false)
	}

	if authQueryType == "playlist/program" {
		response, _ = h.services.GetResponseCodeArchive(userIp)
	}

	if response == 403 {
		logrus.Errorf("response 403 for ip %v uri %v", userIp, Uri)
	}

	c.Status(response)
}

func detectAuthQueryType(uri string) (string, error) {
	splitted := strings.Split(uri, "/")
	for i, val := range splitted {
		if val == "timeshift" {
			return "timeshift", nil
		}
		if val == "playlist" && i < len(splitted)-1 && splitted[i+1] == "program" {
			return "playlist/program", nil
		}
		if val == "streaming" {
			return "streaming", nil
		}
	}
	return "", fmt.Errorf("unexpected auth query type for %v", uri)
}
