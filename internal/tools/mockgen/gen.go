package mock_generator

import _ "github.com/golang/mock/mockgen/model"

// internal - usually adaptors
//go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/auth/mocks/mockedUserStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/auth Store
//go:generate mockgen -package mockedDeviceStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/manager/devices/mock/mockedDeviceStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/devices Store
//go:generate mockgen -package mockedBankAccountStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts/mock/mockedDeviceStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts Store

// external - usually server clients
// wrong example -> //go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/sqlc/auth/user/mock/mockedUserStore.go github.com/Pietroski/SimpleTransactionDemo/adaptors/services/datastore/postgresql/sqlc/auth/user/ Store
