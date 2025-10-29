package complete

import (
	"github.com/gin-gonic/gin"
)

// complete handles the /complete endpoint
func Complete(c *gin.Context) {
	lang := c.Query("lang")
	prefix := c.Query("prefix")
	
	c.JSON(200, gin.H{
		"completions": []string {"print", "split", lang, prefix},
	})
}
