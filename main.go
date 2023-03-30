package main

import (
	"fmt"
	"net/http"
)

func main() {
	apiURL := "https://go-codescanning-githubactions-cloudrun-wsgwmfbvhq-uc.a.run.app/"
	res, _ := http.Get(apiURL)

	fmt.Println(res.StatusCode)             // HTTPレスポンスステータス
	fmt.Println(res.Proto)                  // HTTPプロトコル
	fmt.Println(res.Header["Date"])         // データを取得した日付と時間
	fmt.Println(res.Header["Content-Type"]) // コンテンツのタイプ
	fmt.Println(res.Request.Method)         // GETかPOST
	fmt.Println(res.Request.URL)            // URL
}
