package service

import (
	"context"

	"github.com/brunoluiz/grpc-example/simple"
	"github.com/brunoluiz/grpc-example/simple/generated/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewServer Initiate a GRPC Server
func NewServer() *GRPCServer {
	users := map[string]*simple.User{
		"foo": &simple.User{
			ID:     "foo",
			Name:   "John Doe",
			Active: true,
		},
		"bar": &simple.User{
			ID:     "bar",
			Name:   "Chauffina Carr",
			Active: true,
		},
		"xyz": &simple.User{
			ID:     "xyz",
			Name:   "Pelican Steve",
			Active: true,
		},
	}

	return &GRPCServer{users}
}

// GRPCServer Defines a GRPC Service server
type GRPCServer struct {
	users map[string]*simple.User
}

// GetUser Get user by ID
func (g *GRPCServer) GetUser(_ context.Context, req *api.GetUserRequest) (*api.GetUserResponse, error) {
	user, ok := g.users[req.GetUserId()]
	if !ok {
		return nil, status.Error(codes.NotFound, "resource was not found")
	}

	return &api.GetUserResponse{
		User: &api.User{
			UserId: user.ID,
			Name:   user.Name,
			Active: user.Active,
		},
	}, nil
}

// GetUsers Get all users
func (g *GRPCServer) GetUsers(_ context.Context, _ *api.GetUsersRequest) (*api.GetUsersResponse, error) {
	users := make([]*api.User, 0, len(g.users))
	for _, user := range g.users {
		users = append(users, &api.User{
			UserId: user.ID,
			Name:   user.Name,
			Active: user.Active,
		})
	}

	return &api.GetUsersResponse{
		Users: users,
	}, nil
}
