package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/auth/internal/model"
	"github.com/markgenuine/auth/internal/repository"
	repoMock "github.com/markgenuine/auth/internal/repository/mocks"
	authService "github.com/markgenuine/auth/internal/service/auth_v1"
	"github.com/markgenuine/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type userRepositoryMocksFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx     context.Context
		request *model.UserUpdate
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		repoErr = errors.New("failed updated user by id")

		userID = int64(gofakeit.Uint64())
		name   = gofakeit.Name()
		email  = gofakeit.Email()
		role   = auth_v1.Role_USER.String()

		request = &model.UserUpdate{
			ID:    userID,
			Name:  &name,
			Email: &email,
			Role:  &role,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMocksFunc
	}{
		{
			name: "sucсess",
			args: args{
				ctx:     ctx,
				request: request,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, request).Return(nil)
				return mock
			},
		},
		{
			name: "error",
			args: args{
				ctx:     ctx,
				request: request,
			},
			err: repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, request).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userRepoMock := tt.userRepositoryMock(mc)
			service := authService.NewService(userRepoMock)

			err := service.Update(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
		})
	}
}
