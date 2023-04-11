package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const Api1Url = "https://cloudrun-api1-morikawa-test-wsgwmfbvhq-uc.a.run.app"
const Api2Url = "https://cloudrun-api2-morikawa-test-wsgwmfbvhq-uc.a.run.app"

// 実際は以下構造体はどこかで共通化する想定
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	// router.POST("/users", postUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	router.Run("0.0.0.0:" + port)
}

// TODO: errハンドリング
func getAlbums(c *gin.Context) {
	res, _ := http.Get(Api1Url + "/albums")
	defer res.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))

	c.IndentedJSON(http.StatusOK, string(body))
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	res, _ := http.Get(Api1Url + "/albums/" + string(id))
	defer res.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))

	c.IndentedJSON(http.StatusOK, string(body))
}


func getUsers(c *gin.Context) {
	res, _ := http.Get(Api2Url + "/users")
	defer res.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))

	c.IndentedJSON(http.StatusOK, string(body))
}

func getUserByID(c *gin.Context) {
	id := c.Param("id")
	res, _ := http.Get(Api2Url + "/users/" + string(id))
	defer res.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(res.Body)
	log.Println(string(body))

	c.IndentedJSON(http.StatusOK, string(body))
}
