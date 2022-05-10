package internal_gin_transactions_validators

import (
	"net/http"
	"reflect"
	"testing"

	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
	mocked_validator_v10 "github.com/Pietroski/SimpleTransactionDemo/pkg/mocks/validator-v10"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/mock/gomock"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestCoinCustomValidation(t *testing.T) {
	tests := []struct {
		name   string
		setup  func()
		assert func()
	}{
		{
			name: "invalid coin",
			setup: func() {
				setupValidators(t)
			},
			assert: func() {
				var coin manager_models.CoinHistoryRequest
				req, err := http.NewRequest(
					http.MethodGet, "some-path?coin=NOT_A-COIN", nil,
				)
				require.NoError(t, err)
				ctx := gin.Context{
					Request: req,
				}
				err = ctx.ShouldBindQuery(&coin)
				require.Error(t, err)
			},
		},
		{
			name: "valid coin - pietroski coin",
			setup: func() {
				setupValidators(t)
			},
			assert: func() {
				var coin manager_models.CoinHistoryRequest
				req, err := http.NewRequest(
					http.MethodGet, "some-path?coin=PIETROSKI-COIN", nil,
				)
				require.NoError(t, err)
				ctx := gin.Context{
					Request: req,
				}
				err = ctx.ShouldBindQuery(&coin)
				require.NoError(t, err)
			},
		},
		{
			name: "empty coin with query field",
			setup: func() {
				setupValidators(t)
			},
			assert: func() {
				var coin manager_models.CoinHistoryRequest
				req, err := http.NewRequest(
					http.MethodGet, "some-path?coin=", nil,
				)
				require.NoError(t, err)
				ctx := gin.Context{
					Request: req,
				}
				err = ctx.ShouldBindQuery(&coin)
				require.NoError(t, err)
			},
		},
		{
			name: "empty coin with no query field",
			setup: func() {
				setupValidators(t)
			},
			assert: func() {
				var coin manager_models.CoinHistoryRequest
				req, err := http.NewRequest(
					http.MethodGet, "some-path", nil,
				)
				require.NoError(t, err)
				ctx := gin.Context{
					Request: req,
				}
				err = ctx.ShouldBindQuery(&coin)
				require.NoError(t, err)
			},
		},
		{
			name: "empty coin with no query field",
			setup: func() {
				setupValidators(t)
			},
			assert: func() {
				ctrl := gomock.NewController(t)
				mockedValidator := mocked_validator_v10.NewMockFieldLevelValidator(ctrl)
				num := int64(1234)
				reflectVal := reflect.ValueOf(num)
				mockedValidator.EXPECT().Field().Times(1).Return(reflectVal)

				ok := CoinCustomValidation(mockedValidator)
				require.False(t, ok)
			},
		},
	}
	for _, tt := range tests {
		tt.setup()
		tt.assert()
	}
}

func setupValidators(t *testing.T) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation(
			"CoinCustomValidation",
			CoinCustomValidation,
		)
		require.NoError(t, err)
	}
}
