package main

import (
	"log"
	"net"

	"github.com/markgenuine/auth/internal/app/auth_v1"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	hostGRPC = "localhost:50551"
)

func main() {
	lis, err := net.Listen("tcp", hostGRPC)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, auth_v1.NewAuth())

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
