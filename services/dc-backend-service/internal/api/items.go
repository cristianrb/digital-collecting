package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) getItems(ctx *gin.Context) {
	offsetParam := ctx.Query("offset")
	limitParam := ctx.Query("limit")
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		limit = 100
	}

	println(offset)
	println(limit)
	items, err := s.Storage.GetAllItems(offset, limit)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, items)
}
