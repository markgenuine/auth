package tests

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/auth/internal/api/auth"
	"github.com/markgenuine/auth/internal/converter"
	"github.com/markgenuine/auth/internal/model"
	"github.com/markgenuine/auth/internal/service"
	serviceMock "github.com/markgenuine/auth/internal/service/mocks"
	"github.com/markgenuine/auth/pkg/auth_v1"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type authServiceMocksFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx     context.Context
		request *desc.GetRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")

		userID    = gofakeit.Uint64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		role      = auth_v1.Role_USER
		createdAt = gofakeit.Date()
		updatedAt = gofakeit.Date()

		request = &desc.GetRequest{
			Id: int64(userID),
		}

		outputData = &model.User{
			ID:        int64(userID),
			Name:      name,
			Email:     email,
			Role:      role.String(),
			CreatedAt: createdAt,
			UpdatedAt: sql.NullTime{
				Time:  updatedAt,
				Valid: true,
			},
		}

		response = converter.GetToUserFromService(outputData)
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		authServiceMock authServiceMocksFunc
	}{
		{
			name: "success",
			args: args{
				ctx:     ctx,
				request: request,
			},
			want: response,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, int64(userID)).Return(outputData, nil)
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
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, int64(userID)).Return(nil, serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewImplementation(authServiceMock)

			result, err := api.Get(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
