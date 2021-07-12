package service

import "github.com/unheilbar/hls_frontend_api/pkg/cache"

type UsersCacheService struct {
	cache cache.UsersCache
}

func NewUsersCacheListService(cache cache.UsersCache) *UsersCacheService {
	return &UsersCacheService{cache: cache}
}

func (cs *UsersCacheService) ClearUserCacheByIp(userIp string) {

}
func (cs *UsersCacheService) AddUserCacheItem(item cache.UserCacheItem) {

}
