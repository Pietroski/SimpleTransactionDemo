package mock_generator

import _ "github.com/golang/mock/mockgen/model"

// internal - usually adaptors
//go:generate mockgen -package mockedUserStore -destination ../../../internal/adaptors/datastore/postgresql/auth/sqlc/mocks/mockedUserStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/auth/sqlc Store
//go:generate mockgen -package mockedDeviceStore -destination ../../../internal/adaptors/datastore/postgresql/manager/devices/sqlc/mock/mockedDeviceStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/devices/sqlc Store
//go:generate mockgen -package mockedTransactionStore -destination ../../../internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc/mock/mockedDeviceStore.go github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc Store

// external - usually server clients
//go:generate mockgen -package mocked_gin -destination ../../../pkg/mocks/gin/mocked_response_writer.go github.com/Pietroski/SimpleTransactionDemo/pkg/mocks/gin ResponseWriter
//go:generate mockgen -package mocked_validator_v10 -destination ../../../pkg/mocks/validator-v10/mocked_validator.go github.com/Pietroski/SimpleTransactionDemo/pkg/mocks/validator-v10 FieldLevelValidator
