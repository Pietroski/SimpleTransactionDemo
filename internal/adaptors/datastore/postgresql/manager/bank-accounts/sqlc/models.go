// Code generated by sqlc. DO NOT EDIT.

package sqlc_bank_account_store

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CryptoCurrencies string

const (
	CryptoCurrenciesBITCOIN       CryptoCurrencies = "BITCOIN"
	CryptoCurrenciesDODGECOIN     CryptoCurrencies = "DODGE-COIN"
	CryptoCurrenciesETHEREUM      CryptoCurrencies = "ETHEREUM"
	CryptoCurrenciesPIETROSKICOIN CryptoCurrencies = "PIETROSKI-COIN"
)

func (e *CryptoCurrencies) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CryptoCurrencies(s)
	case string:
		*e = CryptoCurrencies(s)
	default:
		return fmt.Errorf("unsupported scan type for CryptoCurrencies: %T", src)
	}
	return nil
}

type Account struct {
	RowID     int64     `json:"rowID"`
	AccountID uuid.UUID `json:"accountID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type EntryRecord struct {
	RowID     int64            `json:"rowID"`
	AccountID uuid.UUID        `json:"accountID"`
	WalletID  uuid.UUID        `json:"walletID"`
	Coin      CryptoCurrencies `json:"coin"`
	Amount    int64            `json:"amount"`
	CreatedAt time.Time        `json:"createdAt"`
}

type TransactionRecord struct {
	RowID         int64            `json:"rowID"`
	FromAccountID uuid.UUID        `json:"fromAccountID"`
	FromWalletID  uuid.UUID        `json:"fromWalletID"`
	ToAccountID   uuid.UUID        `json:"toAccountID"`
	ToWalletID    uuid.UUID        `json:"toWalletID"`
	Coin          CryptoCurrencies `json:"coin"`
	Amount        int64            `json:"amount"`
	CreatedAt     time.Time        `json:"createdAt"`
}

type Wallet struct {
	RowID     int64            `json:"rowID"`
	WalletID  uuid.UUID        `json:"walletID"`
	AccountID uuid.UUID        `json:"accountID"`
	Coin      CryptoCurrencies `json:"coin"`
	Balance   int64            `json:"balance"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
}
