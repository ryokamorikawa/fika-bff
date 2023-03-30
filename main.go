package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://go-codescanning-githubactions-cloudrun-wsgwmfbvhq-uc.a.run.app/"
	res, _ := http.Get(apiURL)

	fmt.Println(res.StatusCode) // HTTPレスポンスステータス
	fmt.Println(res.Proto)      // HTTPプロトコル
}
