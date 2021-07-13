package service

import (
	"fmt"

	"github.com/unheilbar/hls_frontend_api/pkg/cache"
	"github.com/unheilbar/hls_frontend_api/pkg/whoipapi"
)

type AuthCacheService struct {
	cc cache.ChannelsCache
	uc cache.UsersCache
}

func NewCacheAuth(cc cache.ChannelsCache, uc cache.UsersCache) *AuthCacheService {
	return &AuthCacheService{
		cc: cc,
		uc: uc,
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
	userItem, err := whoipapi.FetchUserItemByIp(userIp)
	if err != nil {
		return 403, fmt.Errorf("error occured during fetching data for %v", userIp)
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
	userItem, err := whoipapi.FetchUserItemByIp(userIp)
	if err != nil {
		return 403, fmt.Errorf("error occured during fetching data for %v", userIp)
	}

	// add user in cache in case of success
	a.uc.AddUserCacheItem(userIp, userItem)

	if userItem.Arh {
		return 200, nil
	} else {
		return 403, nil
	}
}
