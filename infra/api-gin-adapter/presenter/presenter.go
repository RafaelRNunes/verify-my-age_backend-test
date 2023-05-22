package presenter

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/dto"
	"github.com/gin-gonic/gin"
)

func Response[T dto.UserOutput | []dto.UserOutput | string](status bool, data T, err string) *gin.H {
	return &gin.H{
		"status": status,
		"data":   data,
		"error":  err,
	}
}
