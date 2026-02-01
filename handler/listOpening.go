package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauannlima/api_golang/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings from database")
		return
	}
	sendSuccess(ctx, "list-opening", openings)
}
