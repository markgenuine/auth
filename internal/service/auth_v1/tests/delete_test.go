package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/auth/internal/repository"
	repoMock "github.com/markgenuine/auth/internal/repository/mocks"
	authService "github.com/markgenuine/auth/internal/service/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type userRepositoryMocksFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx     context.Context
		request int64
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		repoErr = errors.New("failed to delete user")

		userID  = gofakeit.Uint64()
		request = int64(userID)
	)

	tests := []struct {
		name               string
		args               args
		err                error
		userRepositoryMock userRepositoryMocksFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			err: nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, int64(userID)).Return(nil)
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
				mock.DeleteMock.Expect(ctx, int64(userID)).Return(repoErr)
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

			err := service.Delete(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
		})
	}
}
