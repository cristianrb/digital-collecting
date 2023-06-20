package api

import (
	"dc-backend/pkg/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) getItems(ctx *gin.Context) {
	offset, limit := getOffsetAndLimit(ctx)
	items, err := s.Storage.GetAllItems(offset, limit)
	if err != nil {
		errorResponse(ctx, types.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, items)
}

func (s *Server) buyItem(ctx *gin.Context) {
	accessToken := ctx.Request.Header.Get("Authorization")
	payload, err := s.JWTToken.VerifyToken(accessToken)
	if err != nil {
		errorResponse(ctx, types.ApiError{StatusCode: http.StatusUnauthorized, Message: err.Error()})
		return
	}

	item, apiErr := s.getItemFromDB(ctx, err)
	if err != nil {
		errorResponse(ctx, *apiErr)
		return
	}

	hasItem := s.checkUserHasItem(item.Id, payload.ID)
	if hasItem != nil {
		errorResponse(ctx, *hasItem)
		return
	}

	resp, usersErr := s.UsersClient.RetrieveCoins(accessToken, item.Price)
	if usersErr != nil {
		errorResponse(ctx, *usersErr)
		return
	}

	err = s.Storage.AddItemToUser(item.Id, payload.ID)
	if err != nil {
		errorResponse(ctx, types.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (s *Server) getCollection(ctx *gin.Context) {
	userIdParam := ctx.Param("id")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		errorResponse(ctx, types.ApiError{StatusCode: http.StatusBadRequest, Message: "Invalid user id"})
		return
	}
	offset, limit := getOffsetAndLimit(ctx)

	items, err := s.Storage.GetItemsByUserId(int64(userId), offset, limit)
	if err != nil {
		errorResponse(ctx, types.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, items)
}

func (s *Server) checkUserHasItem(itemId, userId int64) *types.ApiError {
	counter, err := s.Storage.CheckUserHasItem(itemId, userId)
	if err != nil {
		return &types.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()}
	}
	if counter > 0 {
		return &types.ApiError{StatusCode: http.StatusBadRequest, Message: "user already has bought this item"}
	}

	return nil
}

func (s *Server) getItemFromDB(ctx *gin.Context, err error) (*types.Item, *types.ApiError) {
	itemId := types.ItemId{}
	err = ctx.ShouldBindJSON(&itemId)
	if err != nil {
		return nil, &types.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()}
	}
	item, err := s.Storage.GetItemById(itemId.Id)
	if err != nil {
		return nil, &types.ApiError{StatusCode: http.StatusBadRequest, Message: err.Error()}
	}
	return item, nil
}

func getOffsetAndLimit(ctx *gin.Context) (int, int) {
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
	return offset, limit
}
