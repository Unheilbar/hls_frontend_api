package cache

import (
	"sync"

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
	c.ChannelsCache = make(map[string]int, len(chanInfo))
	for channel, chanId := range chanInfo {
		c.ChannelsCache[channel] = chanId.Id
	}
}
