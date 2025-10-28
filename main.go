package main

import (
	// "github.com/gin-gonic/gin"
	// "net/http"
	// "personalized-lsp/complete"
	"log"
	"personalized-lsp/automation"
)


func main() {
	file :=		automation.ListFilesWithPath("..")
	
	for _, f := range file {
			log.Println(f)
	}


	// r := gin.Default()


	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello, World!")
	// })


	// r.GET("complete", complete.Complete)

	// r.Run() 
}
