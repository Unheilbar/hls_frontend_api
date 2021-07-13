package whoipapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/unheilbar/hls_frontend_api/pkg/cache"
)

const (
	baseWhoIpUrl = "http://vladlink.tv/playlist/whocha/whoip/?hlswhoip="
)

type WhoipApiResponse struct {
	Uid int
	Arh int
	Ser []int
}

func GetWhoApiResponse(userIp string) (WhoipApiResponse, error) {
	result := &WhoipApiResponse{}
	err := getJson(baseWhoIpUrl+userIp, result)

	if err != nil {
		return WhoipApiResponse{}, err
	}

	return *result, nil
}

func getJson(url string, target *WhoipApiResponse) error {
	var myClient = &http.Client{Timeout: 3 * time.Second}
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetUserItemFromResponse(r WhoipApiResponse) (cache.UserCacheItem, bool) {
	var access bool
	if r.Arh == 1 {
		access = true
	} else {
		access = false
	}
	item := cache.UserCacheItem{
		Arh:  access,
		Ser:  r.Ser,
		Uid:  r.Uid,
		Time: time.Now().Local(),
	}

	if len(r.Ser) > 0 {
		return item, true
	} else {
		return item, false
	}
}
