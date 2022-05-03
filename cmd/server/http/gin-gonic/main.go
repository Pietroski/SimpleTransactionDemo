package main

import (
	"database/sql"
	"github.com/Pietroski/SimpleTransactionDemo/internal/models/server"
	"log"

	bank_account_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/bank-accounts"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"

	auth_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/auth"
	device_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/manager/devices"
	auth_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/auth"
	manager_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/manager"
)

var (
	stopServerSig = make(chan error)
)

func main() {
	// TODO: pass database conn
	authStore := auth_store.NewStore(&sql.DB{})
	authServer := auth_factory.NewAuthServer(authStore)

	// TODO: pass database conn
	deviceStore := device_store.NewStore(&sql.DB{})
	txStore := bank_account_store.NewStore(&sql.DB{})
	managerServer := manager_factory.NewManagerServer(
		manager_models.Stores{
			DeviceStore: deviceStore,
			TxStore:     txStore,
		},
	)

	//{ // server initialisation grouping
	//
	//	// We do not need waiting group here
	//	// select statement will hold the server running
	//	// until an error is returned during initialisation
	//
	//	// auth server initialisation
	//	go func(stopServerSig chan error) {
	//		if err := authServer.Start(); err != nil {
	//			stopServerSig <- err
	//		}
	//	}(stopServerSig)
	//
	//	// manager server initialisation
	//	go func(stopServerSig chan error) {
	//		if err := managerServer.Start(); err != nil {
	//			stopServerSig <- err
	//		}
	//	}(stopServerSig)
	//}

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
