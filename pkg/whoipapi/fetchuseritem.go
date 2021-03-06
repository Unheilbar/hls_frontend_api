package whoipapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
)

type whoipApiResponse struct {
	Uid int
	Arh int
	Ser []int
}

func FetchUserItemByIp(userIp string, baseWhoIpUrl string, semaphore *chan struct{}) (cache.UserCacheItem, error) {

	result := &whoipApiResponse{}

	err := getJson(baseWhoIpUrl+userIp, result)

	if err != nil {
		<-*semaphore
		return cache.UserCacheItem{}, err
	}
	<-*semaphore
	return getUserItemFromResponse(*result), nil

}

var myClient = &http.Client{Timeout: time.Duration(5) * time.Second}

func getJson(url string, target *whoipApiResponse) error {
	r, err := myClient.Get(url)
	if err != nil {

		logrus.Errorf("Error on address %v, %v", url, err.Error())
		return err
	}

	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(target)

	return err
}

func getUserItemFromResponse(r whoipApiResponse) cache.UserCacheItem {
	var access bool
	if r.Arh == 1 {
		access = true
	} else {
		access = false
	}
	item := cache.UserCacheItem{
		Arh:         access,
		Ser:         r.Ser,
		Uid:         r.Uid,
		CreatedTime: time.Now().Local(),
	}

	return item
}
