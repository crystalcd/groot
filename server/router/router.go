package router

import (
	"github.com/gin-gonic/gin"
)

const Domain = "hackerone.com"

func Rcon(c *gin.Context) {
	
	c.JSON(200, gin.H{"Message": "Hello World"})
}
