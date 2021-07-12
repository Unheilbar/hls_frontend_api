package service

import "github.com/unheilbar/hls_frontend_api/pkg/cache"

type AuthCacheService struct {
	cache cache.UsersCache
}

func NewCacheAuth(cache cache.UsersCache) *AuthCacheService {
	return &AuthCacheService{cache: cache}
}

func (a *AuthCacheService) GetResponseCode(userIp string, channelAllias string) (int, error) {
	return 200, nil
}
