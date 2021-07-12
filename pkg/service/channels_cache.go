package service

import "github.com/unheilbar/hls_frontend_api/pkg/cache"

type ChannelsCacheService struct {
	cache cache.ChannelsCache
}

func NewChannelsCacheService(cache cache.ChannelsCache) *ChannelsCacheService {
	return &ChannelsCacheService{
		cache: cache,
	}
}

func (cs *ChannelsCacheService) ReloadChannels() {

}
