package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauannlima/api_golang/schemas"
)

// @BasePath /api/v1
// @Summary Delete a job opening
// @Description Delete a job opening with the provided id
// @Tags Opening
// @Accept json
// @Produce json
// @Param id query string true "Opening ID"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error deleting opening from database")
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
