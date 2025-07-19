package main

import (
	"log"
	"net"
	"net/http"

	"github.com/asandriss/otp-doc-verification/shared/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Config{}

	if err := config.LoadConfig(&cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := gin.Default()
	router.GET("/health", getHealth)

	ip := "0.0.0.0"
	addr := net.JoinHostPort(ip, cfg.Port)
	router.Run(addr)
}

func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "auth",
	})
}
