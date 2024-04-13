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
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type authServiceMocksFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx     context.Context
		request *desc.UpdateRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")

		userID = gofakeit.Uint64()
		name   = *wrapperspb.String(gofakeit.Name())
		email  = *wrapperspb.String(gofakeit.Email())
		role   = auth_v1.Role_USER

		request = &desc.UpdateRequest{
			Id:    int64(userID),
			Name:  &name,
			Email: &email,
			Role:  role,
		}

		inputData = converter.UpdateToServiceFromUser(request)

		response = &emptypb.Empty{}
	)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
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
				mock.UpdateMock.Expect(ctx, inputData).Return(nil)
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
				mock.UpdateMock.Expect(ctx, inputData).Return(serviceErr)
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

			result, err := api.Update(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
