package service

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
	"github.com/unheilbar/hls_frontend_api/pkg/whoipapi"
)

type AuthCacheService struct {
	cc        cache.ChannelsCache
	uc        cache.UsersCache
	semaphore chan struct{}
}

func NewCacheAuth(cc cache.ChannelsCache, uc cache.UsersCache) *AuthCacheService {
	return &AuthCacheService{
		cc:        cc,
		uc:        uc,
		semaphore: make(chan struct{}, 100), //amount of parallel requests to api
	}
}

func (a *AuthCacheService) GetResponseCodeChannel(userIp string, channelAllias string) (int, error) {
	// check if channel allias exists in cache
	channelId, ok := a.cc.GetChannelId(channelAllias)
	if !ok {
		return 403, fmt.Errorf("channel %v doesn't exist", channelAllias)
	}

	// check if we have user in cache
	userItem, ok := a.uc.GetUserCacheByIp(userIp)
	if ok {
		for _, channel := range userItem.Ser {
			if channelId == channel {
				// user exists in cache and user has channel in his channel list
				return 200, nil
			}
		}
		return 403, fmt.Errorf("user %v has no channel %v in his channel list", userIp, channelAllias)
	}

	// if user doesn't exists then we try to fetch it
	a.semaphore <- struct{}{}
	userItem, err := whoipapi.FetchUserItemByIp(userIp, &a.semaphore)

	// if api response is bad we give access to a user, but do not add user into cache
	if err != nil {
		logrus.Errorf("Bad api response for user %v", userIp)
		return 200, err
	}

	// add user in cache in case of success
	a.uc.AddUserCacheItem(userIp, userItem)

	for _, channel := range userItem.Ser {
		if channel == channelId {
			return 200, nil
		}
	}

	return 403, fmt.Errorf("user %v has no channel %v in his channel list", userIp, channelAllias)
}

func (a *AuthCacheService) GetResponseCodeArchive(userIp string) (int, error) {
	// check if we have user in cache
	userItem, ok := a.uc.GetUserCacheByIp(userIp)
	if ok {
		if userItem.Arh {
			return 200, nil
		} else {
			return 403, nil
		}
	}

	// if user doesn't exists then we try to fetch it
	a.semaphore <- struct{}{}

	userItem, err := whoipapi.FetchUserItemByIp(userIp, &a.semaphore)
	if err != nil {
		logrus.Errorf("error occured when %v data was fetched", err.Error())
		return 200, err
	}

	// add user in cache in case of success
	a.uc.AddUserCacheItem(userIp, userItem)

	if userItem.Arh {
		return 200, nil
	} else {
		return 403, nil
	}
}
