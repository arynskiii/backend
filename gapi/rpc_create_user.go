package gapi

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simple_bank/pb"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
