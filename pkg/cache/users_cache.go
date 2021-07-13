package cache

import (
	"sync"
	"time"

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

func NewUsersCache(expireTime int, cleanupInterval int) *UsersCacheList {
	uc := &UsersCacheList{
		CacheMap:   make(map[string]UserCacheItem, 500),
		expireTime: expireTime,
	}

	uc.StartGC(cleanupInterval)

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
	c.mx.RLock()
	defer c.mx.RUnlock()
	res, ok := c.CacheMap[userIp]
	if ok {
		logrus.Printf("Got user %v from cache ", userIp)
	}

	return res, ok
}

func (c *UsersCacheList) AddUserCacheItem(userIp string, item UserCacheItem) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.CacheMap[userIp] = item
	logrus.Printf("User %v added. Cache size %v  arhv %v", userIp, len(c.CacheMap), item.Arh)
}

func (c *UsersCacheList) ClearUserCacheByIp(userIp string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.CacheMap, userIp)
	logrus.Printf("User with ip %v has been deleted ", userIp)
}

func (c *UsersCacheList) ClearUserCacheByUid(uid int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for key, val := range c.CacheMap {
		if val.Uid == uid {
			delete(c.CacheMap, key)
			logrus.Printf("User with id %v has been deleted ", uid)
		}
	}
}

func (c *UsersCacheList) CleanExpired() {
	c.mx.Lock()
	defer c.mx.Unlock()
	logrus.Printf("GC is starting... User cache size %v ", len(c.CacheMap))
	for key, val := range c.CacheMap {
		if time.Now().Unix()-val.CreatedTime.Unix() > int64(c.expireTime) {
			delete(c.CacheMap, key)
			logrus.Printf("Cache for user id:%v has expired ", key)
		}
	}
	logrus.Printf("GC is finished.. User cache size %v ", len(c.CacheMap))
}
