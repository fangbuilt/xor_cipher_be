package main

import (
	"net/http"
	"xor_cipher/handlers"
	"xor_cipher/middlewares"
	"xor_cipher/services"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "XOR Cipher API is running",
	})
}

func main() {
	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	r.POST("/encrypt", handlers.CreateCipherHandler(services.XOREncrypt))
	r.POST("/decrypt", handlers.CreateCipherHandler(services.XORDecrypt))
	r.GET("/generateKey", handlers.GenerateKeyHandler(services.GenerateKey))
	r.GET("/", healthCheck)

	r.Run(":8080")
}
