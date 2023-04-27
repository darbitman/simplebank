package api

import (
	db "github.com/darbitman/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// Add routes to router.
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}