package cache

import (
	"github.com/unheilbar/hls_frontend_api/pkg/channels_update"
)

type UsersCache interface {
	ClearUserCacheByIp(userIp string)
	ClearUserCacheByUid(uid int)
	AddUserCacheItem(userIp string, item UserCacheItem)
	GetUserCacheByIp(userIp string) (UserCacheItem, bool)
	CleanExpired()
}

type ChannelsCache interface {
	UpdateChannelsCache(map[string]channels_update.ChannelItem)
	GetChannelId(allias string) (int, bool)
}

type Cache struct {
	UsersCache
	ChannelsCache
}

func NewCache() *Cache {
	return &Cache{
		UsersCache:    NewUsersCache(),
		ChannelsCache: NewChannelsCache(),
	}
}
