package cache

type UserCacheItem struct {
}

type UsersCacheList struct {
}

func NewUsersCache() *UsersCacheList {
	return &UsersCacheList{}
}

func (c *UsersCacheList) AddUserCacheItem(item UserCacheItem) {

}

func (c *UsersCacheList) ClearUserCacheByIp(ip string) {

}

func (c *UsersCacheList) GetResponseCode(userIp string, channelAllias string) (int, error) {
	return 200, nil
}
