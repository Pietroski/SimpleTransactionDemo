package internal_gin_custom_validators

import (
	"reflect"
	"testing"

	mocked_field_level_validator "github.com/Pietroski/SimpleTransactionDemo/pkg/mocks/validator-v10"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCoinCustomValidation(t *testing.T) {
	tests := []struct {
		name   string
		setup  func(fl validator.FieldLevel) bool
		stubs  func(fl *mocked_field_level_validator.MockFieldLevelValidator)
		assert func(b bool)
	}{
		{
			name: "",
			setup: func(fl validator.FieldLevel) bool {
				return CoinCustomValidation(fl)
			},
			stubs: func(fl *mocked_field_level_validator.MockFieldLevelValidator) {
				coin := "PIETROSKI-COIN"
				reflectValue := reflect.ValueOf(coin)
				fl.
					EXPECT().
					Field().
					Times(1).
					Return(reflectValue)
			},
			assert: func(b bool) {
				require.Equal(t, true, b)
			},
		},
		{
			name: "",
			setup: func(fl validator.FieldLevel) bool {
				return CoinCustomValidation(fl)
			},
			stubs: func(fl *mocked_field_level_validator.MockFieldLevelValidator) {
				coin := int64(1234)
				reflectValue := reflect.ValueOf(coin)
				fl.
					EXPECT().
					Field().
					Times(1).
					Return(reflectValue)
			},
			assert: func(b bool) {
				require.Equal(t, false, b)
			},
		},
	}
	for _, tt := range tests {
		ctrl := gomock.NewController(t)
		mockedFieldValidator := mocked_field_level_validator.NewMockFieldLevelValidator(ctrl)
		tt.stubs(mockedFieldValidator)
		b := tt.setup(mockedFieldValidator)
		tt.assert(b)
	}
}
