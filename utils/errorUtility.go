package utils

import (
	"github.com/gin-gonic/gin"
)

func Error(message string) gin.H {
	return gin.H{"message": message}
}
