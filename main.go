package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ResponseData はAPIからのレスポンスデータを表す構造体です。
type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	apiURL := "https://go-codescanning-githubactions-cloudrun-wsgwmfbvhq-uc.a.run.app/"
	response, err := callAPI(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("結果:", response)
}

func callAPI(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	return data.Message + "接続できた", nil

}
