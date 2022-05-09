package mocked_auth_middleware

import (
	mocked_gin "github.com/Pietroski/SimpleTransactionDemo/pkg/mocks/gin"
	pkg_auth_extractor "github.com/Pietroski/SimpleTransactionDemo/pkg/tools/extractors/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"reflect"
	"testing"
)

func TestMockedAuthMiddleware(t *testing.T) {
	mockedResponseWriter := mocked_gin.NewMockResponseWriter(gomock.NewController(t))

	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name       string
		args       args
		buildStubs func()
		wantErr    error
	}{
		{
			name: "successfully extracts Bearer token and inject in context",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Header: map[string][]string{
							AuthorizationKey: {"Bearer " + MainMockedBearerToken.String()},
						},
					},
				},
			},
			buildStubs: func() {},
			wantErr:    nil,
		},
		{
			name: "unsuccessfully extracts Bearer token and inject in context",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Header: map[string][]string{
							AuthorizationKey: {"Bearer "},
						},
					},
					Writer: mockedResponseWriter,
				},
			},
			buildStubs: func() {
				mockedResponseWriter.
					EXPECT().
					WriteHeader(gomock.Eq(http.StatusBadRequest)).
					Times(1)
				mockedResponseWriter.
					EXPECT().
					WriteHeaderNow().
					Times(1)
			},
			wantErr: pkg_auth_extractor.ErrInvalidAuthBearerToken,
		},
		{
			name: "unsuccessfully extracts Bearer token and inject in context",
			args: args{
				ctx: &gin.Context{
					Request: &http.Request{
						Header: map[string][]string{
							AuthorizationKey: {"Bearer " + "MainMockedBearerToken.String()"},
						},
					},
					Writer: mockedResponseWriter,
				},
			},
			buildStubs: func() {
				mockedResponseWriter.
					EXPECT().
					WriteHeader(gomock.Eq(http.StatusUnauthorized)).
					Times(1)
				mockedResponseWriter.
					EXPECT().
					WriteHeaderNow().
					Times(1)
			},
			wantErr: pkg_auth_extractor.ErrInvalidAuthBearerToken,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStubs()
			MockedAuthMiddleware(tt.args.ctx)
			t.Log(tt.args.ctx, "||||||||", tt.args.ctx.Errors.Errors(), len(tt.args.ctx.Errors.Errors()))
			if len(tt.args.ctx.Errors.Errors()) == 0 {
				require.NoError(t, tt.wantErr)
				return
			}

			require.Contains(t, tt.wantErr.Error(), tt.args.ctx.Errors.Errors()[0])
		})
	}
}

func TestMockedAuthMiddlewareExtractor(t *testing.T) {
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name    string
		args    args
		want    MockedAuthValues
		wantErr bool
		errDesc error
	}{
		{
			name: "Successfully extracts auth values",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						CtxMockedAuthKey.String(): MockedAuthMap[MainMockedBearerToken],
					},
				},
			},
			want: MockedAuthValues{
				BearerToken: MainMockedBearerToken,
				AccountID:   MockedAuthMap[MainMockedBearerToken].AccountID,
			},
			wantErr: false,
			errDesc: nil,
		},
		{
			name: "Fails to extract auth values",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						"invalid-auth-key": "invalid-Bearer-token",
					},
				},
			},
			want:    MockedAuthValues{},
			wantErr: true,
			errDesc: pkg_auth_extractor.ErrInvalidAuthBearerToken,
		},
		{
			name: "Fails to type assert the token value - wrong string",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						CtxMockedAuthKey.String(): "any-invalid-Bearer-token",
					},
				},
			},
			want:    MockedAuthValues{},
			wantErr: true,
			errDesc: ErrToAssertVar,
		},
		{
			name: "Fails to type assert the token value - wrong type",
			args: args{
				ctx: &gin.Context{
					Keys: map[string]interface{}{
						CtxMockedAuthKey.String(): 0,
					},
				},
			},
			want:    MockedAuthValues{},
			wantErr: true,
			errDesc: ErrToAssertVar,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MockedAuthMiddlewareExtractor(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockedAuthMiddlewareExtractor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockedAuthMiddlewareExtractor() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(err, tt.errDesc) {
				t.Errorf("MockedAuthMiddlewareExtractor() error = %v, errDesc %v", err, tt.errDesc)
			}
		})
	}
}
