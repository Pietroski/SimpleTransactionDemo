package internal_gin_custom_validators

import (
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
	"github.com/go-playground/validator/v10"
)

var CoinCustomValidation validator.Func = func(fl validator.FieldLevel) bool {
	coin, ok := fl.Field().Interface().(manager_models.CryptoCurrencies)
	if ok {
		return coin.IsCryptoCurrency() || coin == ""
	}

	return true
}
