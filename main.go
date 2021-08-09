package main

import (
	"net"
	"os"

	protos "github.com/guilhermelopeseng/api-github-grpc/protos/user"
	"github.com/guilhermelopeseng/api-github-grpc/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	sv := server.NewServer(log)

	protos.RegisterUserServiceServer(gs, sv)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Error("Unable to listen: ", err)
		os.Exit(1)
	}

	gs.Serve(l)

}
