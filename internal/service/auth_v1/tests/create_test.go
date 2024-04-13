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

func TestCreate(t *testing.T) {
	t.Parallel()

	type userRepositoryMocksFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx     context.Context
		request *model.UserCreate
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		repoErr = errors.New("failed to create user")

		userID          = gofakeit.Uint64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(false, true, true, false, false, 6)
		passwordConfirm = password
		role            = auth_v1.Role_USER

		request = &model.UserCreate{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role.String(),
		}
	)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMocksFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: int64(userID),
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, request).Return(int64(userID), nil)
				return mock
			},
		},
		{
			name: "cancel",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, request).Return(0, repoErr)
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

			result, err := service.Create(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
