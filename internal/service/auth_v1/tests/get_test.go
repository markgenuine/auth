package tests

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/auth/internal/converter"
	"github.com/markgenuine/auth/internal/model"
	"github.com/markgenuine/auth/internal/repository"
	modelrepo "github.com/markgenuine/auth/internal/repository/auth/model"
	repoMock "github.com/markgenuine/auth/internal/repository/mocks"
	authService "github.com/markgenuine/auth/internal/service/auth_v1"
	"github.com/markgenuine/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type userRepositoryMocksFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx     context.Context
		request int64
	}

	var (
		ctx     = context.Background()
		mc      = minimock.NewController(t)
		repoErr = errors.New("failed get user by id")

		userID    = int64(gofakeit.Uint64())
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role      = auth_v1.Role_USER.String()
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		request    = userID
		outputData = &modelrepo.User{
			ID:        userID,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		response = converter.GetToServiceFromRepo(outputData)
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMocksFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: response,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(outputData, nil)
				return mock
			},
		},
		{
			name: "error",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMock.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, repoErr)
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

			result, err := service.Get(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
