package health

import "github.com/gin-gonic/gin"

func Check(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "seronz-api hidup",
		"version": "v1.0.0",
	})

}
