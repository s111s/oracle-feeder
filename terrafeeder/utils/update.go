package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"oracle-feeder/terrafeeder/types"
)

// get update price.go data
func UpdatePrices(url string) (*types.History, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_ = resp.Body.Close()

	var history types.History
	err = json.Unmarshal(body, &history)

	return &history, err
}
