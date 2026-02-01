package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauannlima/api_golang/schemas"
)

// @BasePath /api/v1
// @Summary List a job opening
// @Description List all job openings in the database
// @Tags Opening
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings from database")
		return
	}
	sendSuccess(ctx, "list-opening", openings)
}
