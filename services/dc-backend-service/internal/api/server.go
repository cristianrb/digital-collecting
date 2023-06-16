package api

import (
	"dc-backend/internal/storage"

	"github.com/gin-gonic/gin"
)

type Server struct {
	addr    string
	router *gin.Engine
	Storage storage.ItemStorage
}

func New(addr string, storage storage.ItemStorage) *Server {
	server := &Server{
		addr:    addr,
		Storage: storage,
	}

	server.setupRouter()

	return server 
}

func (s *Server) setupRouter() {
	router := gin.Default()
	router.GET("/items", s.getItems)
	s.router = router	
}

func (s *Server) Run() error {
	return s.router.Run(s.addr)
}

func errorResponse(ctx *gin.Context, code int, err error) {
	ctx.JSON(code, gin.H{
		"code":  code,
		"error": err.Error(),
	})
}
