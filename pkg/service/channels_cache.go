package service

import (
	"log"

	"github.com/unheilbar/hls_frontend_api/pkg/cache"
	"github.com/unheilbar/hls_frontend_api/pkg/channels_update"
)

type ChannelsCacheService struct {
	cache cache.ChannelsCache
}

func NewChannelsCacheService(cache cache.ChannelsCache) *ChannelsCacheService {
	return &ChannelsCacheService{
		cache: cache,
	}
}

func (cs *ChannelsCacheService) UpdateChannelsCache() error {
	channelsInfo, err := channels_update.GetChannelsInfo()
	if err != nil {
		log.Fatal("Error occured during updating channels cache")
		return err
	}

	cs.cache.UpdateChannelsCache(channelsInfo)

	return nil
}
