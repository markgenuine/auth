package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	hostGRPC = "localhost:50551"
)

func main() {
	conn, err := grpc.Dial(hostGRPC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed connect to server: %v", err)
	}

	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("failed close connect: %v", err)
		}
	}()

	c := desc.NewUserV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	pass := gofakeit.Password(true, true, true, true, true, 10)

	newUser, err := c.Create(
		ctx,
		&desc.CreateRequest{
			Name:            gofakeit.Name(),
			Email:           gofakeit.Email(),
			Password:        pass,
			PasswordConfirm: pass,
			Role:            0,
		},
	)
	if err != nil {
		log.Printf("failed create user: %v", err)
	}

	fmt.Printf("Create user: %d", newUser.GetId())
}
