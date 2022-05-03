package manager_controller

import (
	device_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager/devices"
	transaction_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager/transactions"
	manager_models "github.com/Pietroski/SimpleTransactionDemo/internal/models/manager"
)

type (
	ManagerController struct {
		Devices *device_controller.DeviceController
		Txs     *transaction_controller.TransactionController
	}
)

func NewManagerController(
	store *manager_models.Stores,
) *ManagerController {
	// TODO: apply validations for arguments if needed

	controller := &ManagerController{
		Devices: device_controller.NewDeviceController(store.DeviceStore),
		Txs:     transaction_controller.NewTransactionController(store.TxStore),
	}

	return controller
}
