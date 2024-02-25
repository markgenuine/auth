package auth_v1

import (
	"context"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Delete user from auth-service
func (s *User) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	fmt.Printf("User delete with ID: %d", req.GetId())

	query, args, err := s.sq.Delete(Users).Where(squirrel.Eq{
		UsersID: req.GetId(),
	}).ToSql()

	if err != nil {
		log.Printf("failed to build query delete user: %s", err.Error())
		return nil, err
	}

	_, err = s.poolDB.Exec(ctx, query, args...)
	if err != nil {
		log.Printf("failed to delete user: %s", err.Error())
		return nil, err
	}

	return &empty.Empty{}, nil
}
