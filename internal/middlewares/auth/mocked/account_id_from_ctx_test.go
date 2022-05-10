package mocked_auth_middleware

import (
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/Pietroski/SimpleTransactionDemo/internal/tools/notification"
	pkg_auth_extractor "github.com/Pietroski/SimpleTransactionDemo/pkg/tools/extractors/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestAccountIdCtxExtractor(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name  string
		args  args
		want  uuid.UUID
		want1 int
		want2 gin.H
	}{
		{
			name: "successfully extracts accountID from context",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						CtxMockedAuthKey.String(): MockedAuthMap[MainMockedBearerToken],
					},
				},
			},
			want:  MainAccountID.ParseForce(),
			want1: 0,
			want2: gin.H{},
		},
		{
			name: "unsuccessfully extracts accountID from context - bad request",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						"invalid-kay": MockedAuthMap[MainMockedBearerToken],
					},
				},
			},
			want:  uuid.UUID{},
			want1: http.StatusBadRequest,
			want2: notification.ClientError.Response(pkg_auth_extractor.ErrInvalidAuthBearerToken),
		},
		{
			name: "unsuccessfully extracts accountID from context - bad request",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						CtxMockedAuthKey.String(): "invalid-token",
					},
				},
			},
			want:  uuid.UUID{},
			want1: http.StatusBadRequest,
			want2: notification.ClientError.Response(ErrToAssertVar),
		},
		{
			name: "unsuccessfully extracts accountID from context - internal server error",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						CtxMockedAuthKey.String(): MockedAuthMap[FailureMockedBearerToken],
					},
				},
			},
			want:  uuid.UUID{},
			want1: http.StatusBadRequest,
			want2: notification.ClientError.Response(errors.New("invalid UUID length: 22")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := AccountIdCtxExtractor(tt.args.ctx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AccountIdCtxExtractor() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AccountIdCtxExtractor() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("AccountIdCtxExtractor() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
