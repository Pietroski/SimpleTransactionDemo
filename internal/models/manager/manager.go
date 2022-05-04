package manager_models

import (
	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	sqlc_device_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/devices/sqlc"
)

type (
	Stores struct {
		DeviceStore sqlc_device_store.Store
		TxStore     sqlc_bank_account_store.Store
	}
)
