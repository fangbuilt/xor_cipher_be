package handlers

import (
	"net/http"
	"xor_cipher/models"
	"xor_cipher/services"

	"github.com/gin-gonic/gin"
)

func CreateCipherHandler(cipherFunc services.CipherFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CipherRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.CipherResponse{
				Success: false,
				Error:   "Invalid request format",
			})
			return
		}

		result, err := cipherFunc(req.Text, req.Key)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.CipherResponse{
				Success: false,
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, models.CipherResponse{
			Success: true,
			Result:  result,
		})
	}
}
