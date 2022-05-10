// Code generated by sqlc. DO NOT EDIT.

package sqlc_bank_account_store

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	GetAccountWallets(ctx context.Context, accountID uuid.UUID) ([]Wallet, error)
	GetPaginatedAccountWallets(ctx context.Context, arg GetPaginatedAccountWalletsParams) ([]Wallet, error)
	GetPaginatedWalletsByAccountID(ctx context.Context, arg GetPaginatedWalletsByAccountIDParams) ([]Wallet, error)
	GetTxWallet(ctx context.Context, arg GetTxWalletParams) (Wallet, error)
	GetWalletsByAccountID(ctx context.Context, accountID uuid.UUID) ([]Wallet, error)
	ListCoinEntryLogsByAccountID(ctx context.Context, arg ListCoinEntryLogsByAccountIDParams) ([]EntryRecord, error)
	ListEntryLogs(ctx context.Context) ([]EntryRecord, error)
	ListEntryLogsByAccountID(ctx context.Context, accountID uuid.UUID) ([]EntryRecord, error)
	ListFromAccountTransactionLogs(ctx context.Context, fromAccountID uuid.UUID) ([]TransactionRecord, error)
	ListPaginatedCoinEntryLogsByAccountID(ctx context.Context, arg ListPaginatedCoinEntryLogsByAccountIDParams) ([]EntryRecord, error)
	ListPaginatedEntryLogs(ctx context.Context, arg ListPaginatedEntryLogsParams) ([]EntryRecord, error)
	ListPaginatedEntryLogsByAccountID(ctx context.Context, arg ListPaginatedEntryLogsByAccountIDParams) ([]EntryRecord, error)
	ListPaginatedFromAccountTransactionLogs(ctx context.Context, arg ListPaginatedFromAccountTransactionLogsParams) ([]TransactionRecord, error)
	ListPaginatedToAccountTransactionLogs(ctx context.Context, arg ListPaginatedToAccountTransactionLogsParams) ([]TransactionRecord, error)
	ListPaginatedTransactionLogs(ctx context.Context, arg ListPaginatedTransactionLogsParams) ([]TransactionRecord, error)
	ListToAccountTransactionLogs(ctx context.Context, toAccountID uuid.UUID) ([]TransactionRecord, error)
	ListTransactionLogs(ctx context.Context) ([]TransactionRecord, error)
	LogEntry(ctx context.Context, arg LogEntryParams) (EntryRecord, error)
	LogTransaction(ctx context.Context, arg LogTransactionParams) (TransactionRecord, error)
	LogTransactionWithEntriesSimplified(ctx context.Context, arg LogTransactionWithEntriesSimplifiedParams) (TransactionRecord, error)
	UpdateAccountWalletBalance(ctx context.Context, arg UpdateAccountWalletBalanceParams) (Wallet, error)
}

var _ Querier = (*Queries)(nil)
