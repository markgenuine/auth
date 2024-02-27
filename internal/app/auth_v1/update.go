package auth_v1

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
)

// Update user of ID
func (s *User) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {
	fmt.Printf("Update user with ID: %d", req.GetId())

	builder := s.sq.Update(users).Where(squirrel.Eq{usersID: req.GetId()})

	var modify bool
	if req.GetName() != nil {
		builder.Set(usersName, req.GetName())
		modify = true
	}

	if req.GetEmail() != nil {
		builder.Set(usersEmail, req.GetEmail())
		modify = true
	}

	if req.GetRole() != desc.Role_UNKNOWN {
		builder.Set(usersRole, req.GetRole().String())
		modify = true
	}

	if modify {
		builder.Set(usersUpdatedAt, time.Now())
		query, args, err := builder.ToSql()
		if err != nil {
			log.Printf("failed to build query update user: %s", err.Error())
			return nil, err
		}

		_, err = s.poolDB.Exec(ctx, query, args...)
		if err != nil {
			log.Printf("failed to update user: %s", err.Error())
			return nil, err
		}
	}

	return &empty.Empty{}, nil
}
