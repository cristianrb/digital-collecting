package api

import (
	"bytes"
	"dc-backend/pkg/types"
	"encoding/json"
	"errors"
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

func (s *Server) buyItem(ctx *gin.Context) {
	itemId := types.ItemId{}
	err := ctx.ShouldBindJSON(&itemId)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	item, err := s.Storage.GetItemById(itemId.Id)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	counter, err := s.Storage.CheckUserHasItem(item.Id, 1)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	if counter > 0 {
		errorResponse(ctx, http.StatusBadRequest, errors.New("user already has bought this item"))
		return
	}

	accessToken := ctx.Request.Header.Get("Authorization")

	coinsToRetrieve := types.Coins{
		Coins: item.Price,
	}
	body, err := json.Marshal(coinsToRetrieve)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req, _ := http.NewRequest("POST", "http://localhost:8080/users/retrieve", bytes.NewBuffer(body))
	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		userResp := types.User{}
		json.NewDecoder(resp.Body).Decode(&userResp)
		println(userResp.Id)

		err = s.Storage.AddItemToUser(item.Id, 1)
		if err != nil {
			errorResponse(ctx, http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, nil)
	} else {
		apiError := types.ApiError{}
		json.NewDecoder(resp.Body).Decode(&apiError)
		errorResponse(ctx, resp.StatusCode, errors.New(apiError.Message))
		return
	}

}
