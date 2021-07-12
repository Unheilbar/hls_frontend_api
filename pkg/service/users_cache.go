package service

import "github.com/unheilbar/hls_frontend_api/pkg/cache"

type UsersCacheService struct {
	cache cache.UsersCache
}

func NewUsersCacheListService(cache cache.UsersCache) *UsersCacheService {
	return &UsersCacheService{cache: cache}
}

func (cs *UsersCacheService) ClearUserCacheByIp(userIp string) {
	cs.cache.ClearUserCacheByIp(userIp)
}

func (cs *UsersCacheService) GetUserCacheByIp(userIp string) (cache.UserCacheItem, bool) {
	UserItem, ok := cs.cache.GetUserCacheByIp(userIp)
	if !ok {
		//
	}
	return UserItem, ok
}

func (cs *UsersCacheService) AddUserCacheItem(userIp string, item cache.UserCacheItem) {
	cs.cache.AddUserCacheItem(userIp, item)
}

func (cs *UsersCacheService) ClearUserCacheByUid(uid int) {
	cs.cache.ClearUserCacheByUid(uid)
}
