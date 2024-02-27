package api_service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiService struct{}

func (apiService ApiService) Index(c *gin.Context) {
	jsonData, err := json.Marshal(map[string]any{
		"/ping": map[string]any{
			"method":   "GET",
			"response": "only response \"pong\" for check connection",
		},
	})
	if err != nil {
		// Handle error appropriately (e.g., log the error and return a 500 Internal Server Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Set the content type header to JSON
	c.Header("Content-Type", "application/json")

	// Write the JSON data to the response
	c.Writer.Write(jsonData)
}

func (apiService ApiService) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
