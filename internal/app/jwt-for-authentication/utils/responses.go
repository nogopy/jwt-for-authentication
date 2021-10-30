package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespondSuccessful(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func RespondNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func RespondBadRequest(ctx *gin.Context, message string) {
	RespondWithError(ctx, http.StatusBadRequest, message)
}

func RespondUnauthorized(ctx *gin.Context) {
	RespondWithError(ctx, http.StatusUnauthorized, "Unauthorized!")
}

func RespondForbidden(ctx *gin.Context) {
	RespondWithError(ctx, http.StatusForbidden, "Forbidden!")
}

func RespondNotFound(ctx *gin.Context) {
	RespondWithError(ctx, http.StatusNotFound, "Not Found!")
}

func RespondInternalError(ctx *gin.Context) {
	RespondWithError(ctx, http.StatusInternalServerError, "Internal server error")
}

func RespondWithError(ctx *gin.Context, statusCode int, message interface{}) {
	ctx.JSON(statusCode, gin.H{
		"error": gin.H{
			"message": message,
			"code":    statusCode,
		},
	})
}
