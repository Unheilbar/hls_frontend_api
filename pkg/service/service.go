package service

import (
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
)

type Auth interface {
	GetResponseCode(userIp string, channelAllias string) (int, error)
}

type UsersCacheList interface {
	ClearUserCacheByIp(userIp string)
	AddUserCacheItem(item cache.UserCacheItem)
}

type ChannelsCache interface {
	ReloadChannels()
}

type Service struct {
	Auth
	UsersCacheList
	ChannelsCache
}

func NewService(cache *cache.Cache) *Service {
	return &Service{
		Auth:           NewCacheAuth(cache.UsersCache),
		UsersCacheList: NewUsersCacheListService(cache.UsersCache),
		ChannelsCache:  NewChannelsCacheService(cache.ChannelsCache),
	}
}
