package http

import (
	"bytes"
	"encoding/json"
	"github.com/jeek120/goutil/log"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		log.L.Errorf("Get %s error: %s", url, err)
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func PostJson(url string, data interface{}, result interface{}) int {
	body, err := json.Marshal(data)
	if err != nil {
		log.L.Errorf("Post %s error: %s", url, err)
		return 500
	}
	resp, err := myClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		log.L.Errorf("Post %s error: %s", url, err)
		return 500
	}
	if result != nil {
		err = json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			log.L.Errorf("Post %s error:序列化结果失败:%s", url, err)
			return 500
		}
	}
	return resp.StatusCode
}
