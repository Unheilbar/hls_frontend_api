package cache

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api/pkg/channels_update"
)

type ChannelsCacheList struct {
	mx            sync.Mutex
	ChannelsCache map[string]int
}

func NewChannelsCache() *ChannelsCacheList {
	return &ChannelsCacheList{
		ChannelsCache: make(map[string]int, 3500),
	}
}

func (c *ChannelsCacheList) UpdateChannelsCache(chanInfo map[string]channels_update.ChannelItem) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.ChannelsCache = make(map[string]int, len(chanInfo))
	for channel, chanId := range chanInfo {
		c.ChannelsCache[channel] = chanId.Id
	}
	logrus.Infof("Channels cache updated. Cache size %v ", len(c.ChannelsCache))
}

func (c *ChannelsCacheList) GetChannelId(allias string) (int, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	id, ok := c.ChannelsCache[allias]
	return id, ok
}
