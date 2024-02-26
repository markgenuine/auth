package auth_v1

import (
	"context"
	"fmt"
	"log"

	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Create user in auth-service
func (s *User) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Printf("Create User: %s", req.GetName())
	fmt.Printf("Create User email: %s", req.GetEmail())
	fmt.Printf("Create User role: %s", req.GetRole())

	query, args, err := s.sq.Insert(users).
		Columns(usersName, usersEmail, usersPassword, usersRole).
		Values(req.GetName(), req.GetEmail(), req.GetPassword(), req.GetRole()).
		Suffix(fmt.Sprintf("RETURNING %s", usersID)).ToSql()

	if err != nil {
		log.Printf("failed to build query create user: %s", err.Error())
		return nil, err
	}

	var userID int64
	err = s.poolDB.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Printf("failed to insert user: %s", err.Error())
		return nil, err
	}

	return &desc.CreateResponse{
		Id: userID,
	}, nil

}
