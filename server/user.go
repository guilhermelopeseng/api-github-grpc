package server

import (
	"context"
	"fmt"
	"net/http"

	protos "github.com/guilhermelopeseng/api-github-grpc/protos/user"
	"github.com/hashicorp/go-hclog"
)

type Server struct {
	log hclog.Logger
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Avatar   string `json:"avatar_url"`
}

func NewServer(log hclog.Logger) *Server {
	return &Server{log: log}
}

func (s *Server) GetUser(ctx context.Context, r *protos.UserRequest) (*protos.UserResponse, error) {
	username := r.GetUsername()
	s.log.Info("Handle GetUser", "User:", username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", username))
	if err != nil {
		s.log.Error("Conection failed with Github:", err)
	}

	usr := &User{}
	err = FromJson(usr, res.Body)
	if err != nil {
		s.log.Error("Reading Failed:", err)
	}

	return &protos.UserResponse{
		Id:   usr.ID,
		Name: usr.Name,
		Info: &protos.Info{
			Username: usr.Username,
			Bio:      usr.Bio,
			Location: usr.Location,
			Avatar:   usr.Avatar,
		},
	}, nil

}
