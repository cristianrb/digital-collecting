package api

import (
	"bytes"
	"dc-backend/pkg/types"
	"encoding/json"
	"net/http"
)

type UsersClient interface {
	RetrieveCoins(accessToken string, coins float64) (*types.User, *types.ApiError)
}
type UsersClientImpl struct {
}

func NewUsersClient() UsersClient {
	return &UsersClientImpl{}
}

func (u *UsersClientImpl) RetrieveCoins(accessToken string, coins float64) (*types.User, *types.ApiError) {
	body, err := json.Marshal(types.Coins{Coins: coins})
	if err != nil {
		return nil, &types.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}

	req, err := http.NewRequest("POST", "http://localhost:8080/users/retrieve", bytes.NewBuffer(body))
	if err != nil {
		return nil, &types.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, &types.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		userResp := types.User{}
		err := json.NewDecoder(resp.Body).Decode(&userResp)
		if err != nil {
			return nil, &types.ApiError{
				StatusCode: http.StatusBadRequest,
				Message:    err.Error(),
			}
		}
		return &userResp, nil
	}

	apiError := &types.ApiError{}
	err = json.NewDecoder(resp.Body).Decode(&apiError)
	if err != nil {
		return nil, &types.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	return nil, apiError
}
