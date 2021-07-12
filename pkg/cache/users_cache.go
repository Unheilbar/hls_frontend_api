package cache

import (
	"sync"
	"time"
)

type UserCacheItem struct {
	Uid  int
	Time time.Time
	Seg  []int
	Arh  bool
}

type UsersCacheList struct {
	mx       sync.RWMutex
	CacheMap map[string]UserCacheItem
}

func NewUsersCache() *UsersCacheList {
	return &UsersCacheList{
		CacheMap: make(map[string]UserCacheItem, 500),
	}
}

func (c *UsersCacheList) GetUserCacheByIp(userIp string) (UserCacheItem, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	res, ok := c.CacheMap[userIp]

	return res, ok
}

func (c *UsersCacheList) AddUserCacheItem(userIp string, item UserCacheItem) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.CacheMap[userIp] = item
}

func (c *UsersCacheList) ClearUserCacheByIp(userIp string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.CacheMap, userIp)
}

func (c *UsersCacheList) ClearUserCacheByUid(uid int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for key, val := range c.CacheMap {
		if val.Uid == uid {
			delete(c.CacheMap, key)
		}
	}
}

func (c *UsersCacheList) GetResponseCode(userIp string, channelAllias string) (int, error) {
	return 200, nil
}
