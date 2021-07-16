package cache

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type UserCacheItem struct {
	Uid         int
	CreatedTime time.Time
	Ser         []int
	Arh         bool
}

type UsersCacheList struct {
	mx         sync.RWMutex
	CacheMap   map[string]UserCacheItem
	expireTime int
}

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func NewUsersCache() *UsersCacheList {

	cacheExpireTime, err := strconv.Atoi(os.Getenv("user_cache_expire_time"))

	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	cacheCleanupInterval, err := strconv.Atoi(os.Getenv("user_cache_cleanup_interval"))

	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	uc := &UsersCacheList{
		CacheMap:   make(map[string]UserCacheItem, 1500),
		expireTime: cacheExpireTime,
	}

	uc.StartGC(cacheCleanupInterval)

	return uc
}

func (c *UsersCacheList) StartGC(cleanupInterval int) {
	go c.GC(cleanupInterval)
}

func (c *UsersCacheList) GC(cleanupInterval int) {
	for {
		<-time.After(time.Duration(cleanupInterval) * time.Second)
		c.CleanExpired()
	}
}

func (c *UsersCacheList) GetUserCacheByIp(userIp string) (UserCacheItem, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	res, ok := c.CacheMap[userIp]
	if ok {
		logrus.Tracef("Got user %v from cache ", userIp)
	}
	return res, ok
}

func (c *UsersCacheList) AddUserCacheItem(userIp string, item UserCacheItem) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.CacheMap[userIp] = item
	logrus.Infof("User %v added. Cache size %v  arhv %v uid %v, time %v", userIp, len(c.CacheMap), item.Arh, item.Uid, item.CreatedTime)
}

func (c *UsersCacheList) ClearUserCacheByIp(userIp string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.CacheMap, userIp)
	logrus.Infof("User with ip %v has been deleted ", userIp)
}

func (c *UsersCacheList) ClearUserCacheByUid(uid int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for key, val := range c.CacheMap {
		if val.Uid == uid {
			delete(c.CacheMap, key)
			logrus.Infof("User with id %v has been deleted ", uid)
		}
	}
}

func (c *UsersCacheList) CleanExpired() {
	c.mx.Lock()
	defer c.mx.Unlock()
	logrus.Infof("GC is starting... User cache size %v ", len(c.CacheMap))
	for key, val := range c.CacheMap {
		// clear cache for users with uid = 0
		if val.Uid == 0 {
			delete(c.CacheMap, key)
			logrus.Infof("Cache for user id:%v has expired ", key)
			continue
		}
		if time.Now().Unix()-val.CreatedTime.Unix() > int64(c.expireTime) {
			delete(c.CacheMap, key)
			logrus.Infof("Cache for user id:%v has expired ", key)
		}
	}
	logrus.Infof("GC is finished.. User cache size %v ", len(c.CacheMap))
}
