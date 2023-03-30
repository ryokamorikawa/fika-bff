package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	apiURL := "https://go-codescanning-githubactions-cloudrun-wsgwmfbvhq-uc.a.run.app/"
	/*接続を試みる*/
	conn, err := net.Dial("tcp", apiURL)
	if err != nil {
		fmt.Println("Error")
	}

	/*GETリクエストを送信*/
	fmt.Fprintf(conn, "GET / HTTP/2.0￥r￥n￥r￥n")

	/*バッファにある返信データをすべて表示*/
	buff := make([]byte, 2048) //ある程度のサイズ
	res, err := bufio.NewReader(conn).Read(buff)
	fmt.Printf("%s", buff[:res])
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
}
