package main

import (
	"github.com/asandriss/otp-doc-verification/auth/internal/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	var cfg config.Config

	err := config.LoadConfig(&cfg)

	router := gin.Default()
	router.GET("/health", getHealth)
	router.Run("localhost:8081")
}

func getHealth(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "{}")
}
