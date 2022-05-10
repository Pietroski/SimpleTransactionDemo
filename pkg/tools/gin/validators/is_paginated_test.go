package pkg_gin_custom_validators

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestIsCorrectlyPaginated(t *testing.T) {
	tests := []struct {
		name       string
		prepareReq func() *gin.Context
		assertResp func(p *Pagination, statusCode int, ginResp gin.H)
	}{
		{
			name: "correct pagination",
			prepareReq: func() *gin.Context {
				rawQuery := fmt.Sprintf("page_size=%v&page_id=%v", 5, 2)
				ctx := &gin.Context{
					Request: &http.Request{
						Method: http.MethodGet,
						URL: &url.URL{
							Host:     "127.0.0.1:80",
							Path:     "/my-happy-path",
							RawQuery: rawQuery,
						},
					},
				}

				return ctx
			},
			assertResp: func(p *Pagination, statusCode int, ginResp gin.H) {
				require.Equal(t, 0, statusCode)
				require.Equal(t, &Pagination{
					Limit:  5,
					Offset: 5,
				}, p)
			},
		},
		{
			name: "correct pagination",
			prepareReq: func() *gin.Context {
				rawQuery := fmt.Sprintf("page_size=%v&page_id=%v", 10, 16)
				ctx := &gin.Context{
					Request: &http.Request{
						Method: http.MethodGet,
						URL: &url.URL{
							Host:     "127.0.0.1:80",
							Path:     "/my-happy-path",
							RawQuery: rawQuery,
						},
					},
				}

				return ctx
			},
			assertResp: func(p *Pagination, statusCode int, ginResp gin.H) {
				require.Equal(t, 0, statusCode)
				require.Equal(t, p, &Pagination{
					Limit:  10,
					Offset: 150,
				})
			},
		},
		{
			name: "fail verification",
			prepareReq: func() *gin.Context {
				rawQuery := fmt.Sprintf("page_size=%v&page_id=%v", 10, 0)
				ctx := &gin.Context{
					Request: &http.Request{
						Method: http.MethodGet,
						URL: &url.URL{
							Host:     "127.0.0.1:80",
							Path:     "/my-happy-path",
							RawQuery: rawQuery,
						},
					},
				}

				return ctx
			},
			assertResp: func(p *Pagination, statusCode int, ginResp gin.H) {
				require.Equal(t, http.StatusBadRequest, statusCode)
				require.Nil(t, p)
			},
		},
		{
			name: "correct pagination",
			prepareReq: func() *gin.Context {
				rawQuery := fmt.Sprintf("page_size=%v", "fail-bind")
				ctx := &gin.Context{
					Request: &http.Request{
						Method: http.MethodGet,
						URL: &url.URL{
							Host:     "127.0.0.1:80",
							Path:     "/my-happy-path",
							RawQuery: rawQuery,
						},
					},
				}

				return ctx
			},
			assertResp: func(p *Pagination, statusCode int, ginResp gin.H) {
				require.Equal(t, http.StatusInternalServerError, statusCode)
				require.Nil(t, p)
			},
		},
		{
			name: "correct pagination",
			prepareReq: func() *gin.Context {
				rawQuery := fmt.Sprintf("page_size=%v&page_id=%v", 0, 0)
				ctx := &gin.Context{
					Request: &http.Request{
						Method: http.MethodGet,
						URL: &url.URL{
							Host:     "127.0.0.1:80",
							Path:     "/my-happy-path",
							RawQuery: rawQuery,
						},
					},
				}

				return ctx
			},
			assertResp: func(p *Pagination, statusCode int, ginResp gin.H) {
				require.Equal(t, 0, statusCode)
				require.Equal(t, p, &Pagination{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := tt.prepareReq
			pagination, statusCode, ginResp := IsCorrectlyPaginated(req())
			tt.assertResp(pagination, statusCode, ginResp)
		})
	}
}

func TestIsPaginated(t *testing.T) {
	type args struct {
		p *Pagination
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is paginated",
			args: args{
				p: &Pagination{
					Limit:  5,
					Offset: 5,
				},
			},
			want: true,
		},
		{
			name: "is paginated",
			args: args{
				p: &Pagination{
					Limit:  0,
					Offset: 0,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPaginated(tt.args.p); got != tt.want {
				t.Errorf("IsPaginated() = %v, want %v", got, tt.want)
			}
		})
	}
}
