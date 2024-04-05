package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/markgenuine/auth/internal/api/auth"
	"github.com/markgenuine/auth/internal/service"
	serviceMock "github.com/markgenuine/auth/internal/service/mocks"
	desc "github.com/markgenuine/auth/pkg/auth_v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type authServiceMocksFunc func(mc *minimock.Controller) service.AuthService

	type args struct {
		ctx     context.Context
		request *desc.DeleteRequest
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		serviceErr = fmt.Errorf("service error")

		userID  = gofakeit.Uint64()
		request = &desc.DeleteRequest{
			Id: int64(userID),
		}

		response = &emptypb.Empty{}
	)

	t.Cleanup(mc.Finish)

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
			want: nil,
			err:  serviceErr,
			authServiceMock: func(mc *minimock.Controller) service.AuthService {
				mock := serviceMock.NewAuthServiceMock(mc)
				mock.DeleteMock.Expect(ctx, int64(userID)).Return(serviceErr)
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

			newID, err := api.Delete(tt.args.ctx, tt.args.request)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
