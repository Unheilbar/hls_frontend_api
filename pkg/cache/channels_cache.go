package cache

import "github.com/unheilbar/hls_frontend_api/pkg/channels_update"

type ChannelsCacheList struct {
}

func NewChannelsCache() *ChannelsCacheList {
	return &ChannelsCacheList{}
}

func (c *ChannelsCacheList) UpdateChannelsCache() {
	channels_update.GetChannelsInfoResponse()
}
