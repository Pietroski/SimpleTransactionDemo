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
		ToAccountID uuid.UUID        `json:"toAccountId"`
		Amount      int64            `json:"amount"`
		Coin        CryptoCurrencies `json:"coin"`
	}

	DepositRequest struct {
		Amount int64            `json:"amount"`
		Coin   CryptoCurrencies `json:"coin"`
	}

	WithdrawRequest struct {
		Amount int64            `json:"amount"`
		Coin   CryptoCurrencies `json:"coin"`
	}

	BalanceRequest struct {
		Coin CryptoCurrencies `json:"coin"`
	}
)
