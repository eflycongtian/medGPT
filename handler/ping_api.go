package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "ping success!",
	})
}
