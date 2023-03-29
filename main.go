package main

import (
	"fmt"
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
	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Println("Error Request:", err)
		return
	}
	fmt.Println("接続できた", response)

}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "BFF"
	}
	fmt.Fprintf(w, "Hello %s!\n", name)
}
