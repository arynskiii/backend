package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "simple_bank/db/sqlc"
	"simple_bank/token"
	"simple_bank/util"
)

type Server struct {
	store      *db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot craete token maker : %w", err)
	}
	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users/", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/auth").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.POST("/accounts/:id", server.deleteAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router

}

func (s *Server) Run(address string) error {
	return s.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
