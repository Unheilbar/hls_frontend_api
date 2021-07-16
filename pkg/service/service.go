package service

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
)

type Auth interface {
	GetResponseCodeChannel(userIp string, channelAllias string, isTimeshift bool) (int, error)
	GetResponseCodeArchive(userIp string) (int, error)
}

type UsersCacheList interface {
	ClearUserCacheByIp(userIp string)
	AddUserCacheItem(userIp string, item cache.UserCacheItem)
	ClearUserCacheByUid(uid int)
	GetUserCacheByIp(userIp string) (cache.UserCacheItem, bool)
}

type ChannelsCache interface {
	UpdateChannelsCache() error
}

type Service struct {
	Auth
	UsersCacheList
	ChannelsCache
}

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func NewService(cache *cache.Cache) *Service {
	whoipApiUrl := os.Getenv("who_ip_url")

	concurrentRequestsLimit, err := strconv.Atoi(os.Getenv("limit_concurrent_api_requests"))

	if err != nil {
		concurrentRequestsLimit = 100
	}

	deps := Dependencies{
		whoipApiUrl:             whoipApiUrl,
		concurrentRequestsLimit: concurrentRequestsLimit,
	}

	return &Service{
		Auth:           NewCacheAuth(cache.ChannelsCache, cache.UsersCache, deps),
		UsersCacheList: NewUsersCacheListService(cache.UsersCache),
		ChannelsCache:  NewChannelsCacheService(cache.ChannelsCache),
	}
}
