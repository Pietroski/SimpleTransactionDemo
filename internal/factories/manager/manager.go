package manager_factory

import (
	"github.com/gin-gonic/gin"

	manager_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager"
	devices_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/manager/devices"
	transaction_factory "github.com/Pietroski/SimpleTransactionDemo/internal/factories/manager/transactions"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
)

type ManagerServer struct {
	address string
	Router  *gin.Engine

	stores manager_models.Stores
}

func NewManagerServer(
	stores manager_models.Stores,
) *ManagerServer {
	// TODO: apply validations for arguments if needed

	factory := &ManagerServer{
		stores: manager_models.Stores{
			DeviceStore: stores.DeviceStore,
			TxStore:     stores.TxStore,
		},
	}

	factory.Handle()

	return factory
}

func (f *ManagerServer) Handle() {
	r := gin.Default()

	// TODO: implement error handling for CORS
	_ = r.SetTrustedProxies([]string{"localhost", "127.0.0.1"})

	managerController := manager_controller.NewManagerController(
		&manager_models.Stores{
			DeviceStore: f.stores.DeviceStore,
			TxStore:     f.stores.TxStore,
		},
	)
	deviceFactory := devices_factory.NewDeviceFactory(managerController.Devices)
	transactionFactory := transaction_factory.NewTransactionFactory(managerController.Txs)

	v1 := r.Group("/v1")
	{
		manager := v1.Group("/manager")
		{
			deviceFactory.Handle(manager)
			transactionFactory.Handle(manager)
		}
	}

	f.Router = r
}

func (f *ManagerServer) Start() error {
	// TODO: remove this after implemented via envs
	f.address = ":8089"
	return f.Router.Run(f.address)
}
