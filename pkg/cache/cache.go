package cache

type UsersCache interface {
	GetResponseCode(userIp string, channelAllias string) (int, error)
	ClearUserCacheByIp(userIp string)
	ClearUserCacheByUid(uid int)
	AddUserCacheItem(userIp string, item UserCacheItem)
	GetUserCacheByIp(userIp string) (UserCacheItem, bool)
}

type ChannelsCache interface {
	UpdateChannelsCache()
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
