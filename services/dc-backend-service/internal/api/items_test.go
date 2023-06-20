package api

import (
	mockdb "dc-backend/mocks/db"
	mocktoken "dc-backend/mocks/token"
	"dc-backend/pkg/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	tokenValidator := mocktoken.NewMockJWTValidator(ctrl)
	storage := mockdb.NewMockItemStorage(ctrl)
	server := New(":8080", tokenValidator, storage)
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/items", nil)

	items := types.Items{
		types.Item{
			Id: 1,
		},
		types.Item{
			Id: 2,
		},
		types.Item{
			Id: 3,
		},
	}
	storage.EXPECT().GetAllItems(0, 100).Times(1).Return(items, nil)

	server.router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
