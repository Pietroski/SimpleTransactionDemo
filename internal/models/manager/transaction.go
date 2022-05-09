package manager_models

import "github.com/google/uuid"

type CryptoCurrencies string

const (
	CryptoCurrenciesBITCOIN       CryptoCurrencies = "BITCOIN"
	CryptoCurrenciesDODGECOIN     CryptoCurrencies = "DODGE-COIN"
	CryptoCurrenciesETHEREUM      CryptoCurrencies = "ETHEREUM"
	CryptoCurrenciesPIETROSKICOIN CryptoCurrencies = "PIETROSKI-COIN"
)

func (cc CryptoCurrencies) String() string {
	return string(cc)
}

type (
	TransactionRequest struct {
		ToAccountID uuid.UUID        `json:"toAccountId" binding:"required"`
		Amount      int64            `json:"amount" binding:"required"`
		Coin        CryptoCurrencies `json:"coin" binding:"required"`
	}

	DepositRequest struct {
		Amount int64            `json:"amount" binding:"required"`
		Coin   CryptoCurrencies `json:"coin" binding:"required"`
	}

	WithdrawRequest struct {
		Amount int64            `json:"amount" binding:"required"`
		Coin   CryptoCurrencies `json:"coin" binding:"required"`
	}

	BalanceRequest struct {
		Coin CryptoCurrencies `json:"coin" binding:"required"`
	}
)
