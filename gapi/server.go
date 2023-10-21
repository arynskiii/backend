package gapi

import (
	"fmt"
	db "simple_bank/db/sqlc"
	"simple_bank/pb"
	"simple_bank/token"
	"simple_bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	store      *db.Store
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

	return server, nil
}
