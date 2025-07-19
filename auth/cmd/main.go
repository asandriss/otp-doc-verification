package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/health", getHealth)

	router.Run("localhost:8081")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "{}")
}
