package cache

import "github.com/unheilbar/hls_frontend_api/pkg/channels_update"

type UsersCache interface {
	GetResponseCode(userIp string, channelAllias string) (int, error)
	ClearUserCacheByIp(userIp string)
	ClearUserCacheByUid(uid int)
	AddUserCacheItem(userIp string, item UserCacheItem)
	GetUserCacheByIp(userIp string) (UserCacheItem, bool)
}

type ChannelsCache interface {
	UpdateChannelsCache(map[string]channels_update.ChannelItem)
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
