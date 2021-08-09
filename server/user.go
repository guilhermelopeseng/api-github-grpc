package server

import (
	"context"
	"fmt"
	"net/http"

	proto "github.com/guilhermelopeseng/api-github-grpc/protos/user"
	"github.com/hashicorp/go-hclog"
)

type Server struct {
	log hclog.Logger
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Avatar   string `json:"avatar_url"`
}

func NewServer(l hclog.Logger) *Server {
	return &Server{l}
}

func (s *Server) GetUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	username := req.GetUsername()
	s.log.Info("Handle GetUser", "User", username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", username))
	if err != nil {
		s.log.Error("Conection failed with Github", err)
		return &proto.UserResponse{}, err
	}

	usr := &User{}
	err = FromJson(usr, res.Body)
	if err != nil {
		s.log.Error("Reading failed", err)
		return &proto.UserResponse{}, err
	}

	return &proto.UserResponse{
		Id:   usr.ID,
		Name: usr.Name,
		Info: &proto.Info{
			Bio:      usr.Bio,
			Location: usr.Location,
			Avatar:   usr.Avatar,
		},
	}, nil
}
