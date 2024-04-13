package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/auth/internal/api/auth"
	"github.com/markgenuine/auth/internal/converter"
	"github.com/markgenuine/auth/internal/service"
	serviceMock "github.com/markgenuine/auth/internal/service/mocks"
	"github.com/markgenuine/auth/pkg/auth_v1"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type authServiceMocksFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx     context.Context
		request *desc.CreateRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")

		userID          = gofakeit.Uint64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		password        = gofakeit.Password(false, true, true, false, false, 6)
		passwordConfirm = password
		role            = auth_v1.Role_USER

		request = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: passwordConfirm,
			Role:            role,
		}

		inputData = converter.CreateUserToServiceFromUser(request)

		response = &auth_v1.CreateResponse{Id: int64(userID)}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
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
				mock.CreateMock.Expect(ctx, inputData).Return(int64(userID), nil)
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
				mock.CreateMock.Expect(ctx, inputData).Return(0, serviceErr)
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

			result, err := api.Create(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
