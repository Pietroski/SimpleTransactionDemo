package manager_models

import (
	bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts"
	device_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/devices"
)

type (
	Stores struct {
		DeviceStore device_store.Store
		TxStore     bank_account_store.Store
	}
)
