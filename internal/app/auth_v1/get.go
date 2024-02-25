package auth_v1

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Get user of ID
func (s *User) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	fmt.Printf("Get user with ID: %d", req.GetId())

	query, args, err := s.sq.Select(
		UsersID, UsersName, UsersEmail, UsersPassword, UsersPasswordConfirm,
		UsersRole, UsersCreatedAt, UsersUpdatedAt).
		Where(squirrel.Eq{UsersID: req.GetId()}).
		ToSql()
	
	if err != nil {
		log.Printf("failed to build query get user: %s", err.Error())
		return nil, err
	}

	var res *desc.GetResponse
	err = s.poolDB.QueryRow(ctx, query, args...).Scan(res)
	if err != nil {
		log.Printf("failed to get user: %s", err.Error())
		return nil, err
	}

	return res, nil
}
