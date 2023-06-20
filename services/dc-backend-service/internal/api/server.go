package api

import (
	"dc-backend/internal/storage"
	"dc-backend/internal/token"
	"dc-backend/pkg/types"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr        string
	router      *gin.Engine
	Storage     storage.ItemStorage
	UsersClient UsersClient
	JWTToken    token.JWTValidator
}

func New(addr string, jwtValidator token.JWTValidator, storage storage.ItemStorage) *Server {
	server := &Server{
		addr:        addr,
		Storage:     storage,
		UsersClient: NewUsersClient(),
		JWTToken:    jwtValidator,
	}

	server.setupRouter()

	return server
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.GET("/items", s.getItems)
	router.GET("/items/:id", s.getCollection)
	router.GET("/profile", s.getProfile)
	router.POST("/items", s.buyItem)
	s.router = router
}

func (s *Server) Run() error {
	return s.router.Run(s.addr)
}

func errorResponse(ctx *gin.Context, error types.ApiError) {
	ctx.JSON(error.StatusCode, error)
}
