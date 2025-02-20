// Package main contains rest server.
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vishal-uttamchandani/go-gold/internal/config"
)

func main() {
	cfg, err := config.FromEnvVariables()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(fmt.Sprintf(":%d", cfg.APIPort)); err != nil {
		panic(err)
	}
}
