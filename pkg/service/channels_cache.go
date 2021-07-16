package service

import (
	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
	"github.com/unheilbar/hls_frontend_api/pkg/channels_update"
)

type ChannelsCacheService struct {
	cache cache.ChannelsCache
}

func NewChannelsCacheService(cache cache.ChannelsCache) *ChannelsCacheService {
	ccs := &ChannelsCacheService{
		cache: cache,
	}
	defer ccs.UpdateChannelsCache()

	return ccs
}

func (cs *ChannelsCacheService) UpdateChannelsCache() error {
	channelsInfo, err := channels_update.GetChannelsInfo()
	if err != nil {
		logrus.Errorf("Error occured during updating channels cache")
		return err
	}

	cs.cache.UpdateChannelsCache(channelsInfo)
	return nil
}
