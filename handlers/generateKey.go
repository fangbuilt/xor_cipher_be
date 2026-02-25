package handlers

import (
	"net/http"
	"xor_cipher/models"
	"xor_cipher/services"

	"github.com/gin-gonic/gin"
)

func GenerateKeyHandler(generateKeyFunc services.GenerateKeyFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := generateKeyFunc()

		c.JSON(http.StatusOK, models.GenerateKeyResponse{
			Result: result,
		})
	}
}
