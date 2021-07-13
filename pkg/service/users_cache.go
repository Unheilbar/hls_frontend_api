package service

import "github.com/unheilbar/hls_frontend_api/pkg/cache"

type UsersCacheService struct {
	cache cache.UsersCache
}

func NewUsersCacheListService(cache cache.UsersCache) *UsersCacheService {
	return &UsersCacheService{cache: cache}
}

func (c *UsersCacheService) ClearUserCacheByIp(userIp string) {
	c.cache.ClearUserCacheByIp(userIp)
}
func (c *UsersCacheService) AddUserCacheItem(userIp string, item cache.UserCacheItem) {
	c.cache.AddUserCacheItem(userIp, item)
}

func (c *UsersCacheService) ClearUserCacheByUid(uid int) {
	c.cache.ClearUserCacheByUid(uid)
}

func (c *UsersCacheService) GetUserCacheByIp(userIp string) (cache.UserCacheItem, bool) {
	return c.cache.GetUserCacheByIp(userIp)
}
