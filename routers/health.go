package routers

import (
	"github.com/gin-gonic/gin"
)

func HealthGET(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "Hi I am UP from TechZone Automation", "cluster": "Edge",
	})
}
