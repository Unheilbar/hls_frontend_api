package channels_update

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ChannelItem struct {
	Id int `json: "id"`
}

func Init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func GetChannelsInfo() (map[string]ChannelItem, error) {

	baseUpdateChannelsUrl := os.Getenv("channels_update_url")

	var result map[string]ChannelItem

	var myClient = &http.Client{Timeout: 3 * time.Second}

	r, err := myClient.Get(baseUpdateChannelsUrl)

	if err != nil {
		return result, err
	}

	resp, errResp := ioutil.ReadAll(r.Body)

	if errResp != nil {
		return result, err
	}

	err = json.Unmarshal(resp, &result)

	return result, err
}
