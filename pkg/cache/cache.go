package cache

type UsersCache interface {
	GetResponseCode(userIp string, channelAllias string) (int, error)
	ClearUserCacheByIp(userIp string)
	AddUserCacheItem(item UserCacheItem)
}

type ChannelsCache interface {
	ReloadChannels()
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
