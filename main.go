package main

import (
	apiService "sss/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api_service := apiService.ApiService{}
	r.GET("/", api_service.Index)
	r.GET("/ping", api_service.Ping)
	r.Run()
}
