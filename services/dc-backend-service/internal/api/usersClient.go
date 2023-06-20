package api

import (
	"bytes"
	"dc-backend/pkg/types"
	"encoding/json"
	"net/http"
)

type UsersClient interface {
	RetrieveCoins(accessToken string, coins float64) (*types.User, *types.ApiError)
	UserInfo(accessToken string) (*types.User, *types.ApiError)
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

	userResp := &types.User{}
	apiError := executeRequest(req, userResp)
	if apiError != nil {
		return nil, apiError
	}
	return userResp, nil
}

func (u *UsersClientImpl) UserInfo(accessToken string) (*types.User, *types.ApiError) {
	req, err := http.NewRequest("GET", "http://localhost:8080/users/me", nil)
	if err != nil {
		return nil, &types.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	req.Header.Set("Authorization", accessToken)
	req.Header.Set("Content-Type", "application/json")

	userResp := &types.User{}
	apiError := executeRequest(req, userResp)
	if apiError != nil {
		return nil, apiError
	}
	return userResp, nil
}

func executeRequest(req *http.Request, payload any) *types.ApiError {
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return &types.ApiError{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		}
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		err := json.NewDecoder(resp.Body).Decode(&payload)
		if err != nil {
			return &types.ApiError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			}
		}
		return nil
	}

	apiError := &types.ApiError{}
	err = json.NewDecoder(resp.Body).Decode(&apiError)
	if err != nil {
		return &types.ApiError{
			StatusCode: apiError.StatusCode,
			Message:    err.Error(),
		}
	}
	return apiError
}
