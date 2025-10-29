package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personalized-lsp/complete"
	"personalized-lsp/data"
)


func main() {
	r := gin.Default()


	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})


	r.GET("complete", complete.Complete)

	r.GET("load_data", data.LoadData)

	r.Run() 
}
