package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// ResponseData はAPIからのレスポンスデータを表す構造体です。
type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	apiURL := "https://go-codescanning-githubactions-cloudrun-wsgwmfbvhq-uc.a.run.app/"
	response, err := callAPI(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("結果:", response)
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
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
