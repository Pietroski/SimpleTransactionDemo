package internal_gin_transactions_validators

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Validator interface {
	Register()
}

type TxValidators struct {
	errRegistryChan      chan error
	CoinCustomValidation validator.Func
}

func NewTxValidator(errRegistryChan chan error) *TxValidators {
	return &TxValidators{
		errRegistryChan:      errRegistryChan,
		CoinCustomValidation: CoinCustomValidation,
	}
}

func (v *TxValidators) Register() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.errRegistryChan <- fmt.Errorf(
			"error to register tx validation: %v",
			val.RegisterValidation(
				"CoinCustomValidation",
				v.CoinCustomValidation,
			))
	}
}
