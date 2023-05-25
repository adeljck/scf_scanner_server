package Modules

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IPInfo struct {
	Country    string `json:"country"`
	RegionName string `json:"regionName"`
	City       string `json:"city"`
	ISP        string `json:"isp"`
	Query      string `json:"query"`
}

func (I *IPInfo) GetCurrentRequestIPInfo() {
	resp, _ := http.Get("http://myexternalip.com/raw")
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	url := "http://ip-api.com/json/" + string(content) + "?lang=zh-CN"
	resp, _ = http.Get(url)
	defer resp.Body.Close()
	content, _ = ioutil.ReadAll(resp.Body)
	json.Unmarshal(content, &I)
}
