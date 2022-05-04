package main

import (
	"database/sql"

	sqlc_auth_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/auth/sqlc"
	sqlc_bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/bank-accounts/sqlc"
	sqlc_device_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/manager/devices/sqlc"

	"log"

	"github.com/Pietroski/SimpleTransactionDemo/internal/models/server"

	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"

	auth_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/auth"
	manager_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/manager"
)

var (
	stopServerSig = make(chan error)
)

func main() {
	// TODO: pass database conn
	authStore := sqlc_auth_store.NewStore(&sql.DB{})
	authServer := auth_factory.NewAuthServer(authStore)

	// TODO: pass database conn
	deviceStore := sqlc_device_store.NewStore(&sql.DB{})
	txStore := sqlc_bank_account_store.NewStore(&sql.DB{})
	managerServer := manager_factory.NewManagerServer(
		manager_models.Stores{
			DeviceStore: deviceStore,
			TxStore:     txStore,
		},
	)

	startServers(
		stopServerSig,
		authServer,
		managerServer,
	)

	select {
	case err := <-stopServerSig:
		log.Println(err)
		return
	}
}

func startServers(stopServerSig chan error, servers ...server.Server) {
	for _, s := range servers {
		go func(stopServerSig chan error, s server.Server) {
			if err := s.Start(); err != nil {
				stopServerSig <- err
			}
		}(stopServerSig, s)
	}
}
