package channels_update

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseUpdateChannelsUrl = "http://vladlink.tv/playlist/getfronthls"
)

type ChannelItem struct {
	Id int `json: "id"`
}

func GetChannelsInfo() (map[string]ChannelItem, error) {
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
