package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)
func main() {
	engine := gin.Default()
	engine.Any("/", WebRoot)
	engine.Run(":8000")
}

func WebRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

