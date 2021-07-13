package whoipapi

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/unheilbar/hls_frontend_api/pkg/cache"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

type whoipApiResponse struct {
	Uid int
	Arh int
	Ser []int
}

func FetchUserItemByIp(userIp string) (cache.UserCacheItem, error) {
	baseWhoIpUrl := os.Getenv("who_ip_url")

	result := &whoipApiResponse{}
	err := getJson(baseWhoIpUrl+userIp, result)

	if err != nil {
		return cache.UserCacheItem{}, err
	}

	return getUserItemFromResponse(*result), nil

}
func getJson(url string, target *whoipApiResponse) error {
	clienTimeout, err := strconv.Atoi(os.Getenv("who_ip_timeout"))

	if err != nil {
		clienTimeout = 5
	}

	var myClient = &http.Client{Timeout: time.Duration(clienTimeout) * time.Second}

	r, err := myClient.Get(url)

	if err != nil {
		logrus.Errorf("Error occured on %v request %v", err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
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
