package mock_generator

import _ "github.com/golang/mock/mockgen/model"

// internal - usually adaptors
//go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/auth/sqlc/mocks/mockedUserStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/auth/sqlc Store
//go:generate mockgen -package mockedDeviceStore -destination ../../../internal/adaptors/datastore/postgresql/manager/devices/sqlc/mock/mockedDeviceStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/devices/sqlc Store
//go:generate mockgen -package mockedTransactionStore -destination ../../../internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc/mock/mockedDeviceStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc Store

// external - usually server clients
// wrong example -> //go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/go-migrate/auth/user/mock/mockedUserStore.go github.com/Pietroski/SimpleTransactionDemo/adaptors/services/datastore/postgresql/go-migrate/auth/user/ Store
