package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	executions "cloud.google.com/go/workflows/executions/apiv1"
	"cloud.google.com/go/workflows/executions/apiv1/executionspb"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

const Api1Url = "https://cloudrun-api1-morikawa-test-wsgwmfbvhq-uc.a.run.app"
const Api2Url = "https://cloudrun-api2-morikawa-test-wsgwmfbvhq-uc.a.run.app"
const WorkflowUrl = "https://workflowexecutions.googleapis.com/v1/projects/kaigofika-poc01/locations/us-central1/workflows/workflow-1-morikawa-test/executions"
const ProjectId = "kaigofika-poc01"
const Location = "us-central1"

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	router.POST("/workflow", executeWorkFlow)

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

func executeWorkFlow(c *gin.Context) {

	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	client, err := executions.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
		fmt.Printf("executions.NewClient: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer client.Close()

	req := &executionspb.CreateExecutionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/workflows/executions/apiv1beta/executionspb#CreateExecutionRequest.
		Parent: "projects/" + ProjectId + "/locations/" + Location + "/workflows/" + "workflow-1-morikawa-test",
	}
	resp, err := client.CreateExecution(ctx, req)
	if err != nil {
		// TODO: Handle error.
		fmt.Printf("client.CreateExecution: %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.Println(resp)

	c.JSON(http.StatusOK, resp)
}
