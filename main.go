package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const Api1Url = "https://cloudrun-api1-morikawa-test-wsgwmfbvhq-uc.a.run.app"
const Api2Url = "https://cloudrun-api2-morikawa-test-wsgwmfbvhq-uc.a.run.app"
const WorkflowUrl = "https://workflowexecutions.googleapis.com/v1/projects/kaigofika-poc01/locations/us-central1/workflows/workflow-1-morikawa-test/executions"

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	router.POST("/workflow", callWorkFlow)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	router.Run("0.0.0.0:" + port)
}

func postUsers(c *gin.Context) {
	fmt.Println("### postUsers")
}

func getAlbums(c *gin.Context) {
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, Api1Url)
	if err != nil {
		fmt.Printf("idtoken.NewClient: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp, err := client.Get(Api1Url + "/albums")
	if err != nil {
		fmt.Printf("client.Get: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))

	c.JSON(resp.StatusCode, string(body))

}

func getAlbumByID(c *gin.Context) {

	id := c.Param("id")
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, Api1Url)
	if err != nil {
		fmt.Printf("idtoken.NewClient: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp, err := client.Get(Api1Url + "/albums/" + string(id))
	if err != nil {
		fmt.Printf("client.Get: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))

	c.JSON(resp.StatusCode, string(body))

}

func getUsers(c *gin.Context) {

	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, Api2Url)
	if err != nil {
		fmt.Printf("idtoken.NewClient: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp, err := client.Get(Api2Url + "/users")
	if err != nil {
		fmt.Printf("client.Get: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))

	c.JSON(resp.StatusCode, string(body))

}

func getUserByID(c *gin.Context) {

	id := c.Param("id")
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, Api2Url)
	if err != nil {
		fmt.Printf("idtoken.NewClient: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp, err := client.Get(Api2Url + "/users/" + string(id))
	if err != nil {
		fmt.Printf("client.Get: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 取得したURLの内容を読み込む
	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))

	c.JSON(resp.StatusCode, string(body))

}

const (
	method      = "POST"
	contentType = "application/json"
)

func callWorkFlow(c *gin.Context) {

	var (
		body = []byte("{}")
		buf  = bytes.NewBuffer(body)
	)

	req, err := http.NewRequest(method, WorkflowUrl, buf)
	if err != nil {
		fmt.Printf("http.NewRequest: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	req.Header.Add("Content-Type", contentType)

	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, WorkflowUrl)
	if err != nil {
		fmt.Printf("idtoken.NewClient: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client.Get: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// 取得したURLの内容を読み込む
	decoder := json.NewDecoder(resp.Body)
	log.Println(decoder)

	c.JSON(resp.StatusCode, decoder)
}
