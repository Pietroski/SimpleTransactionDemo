package devices_factory

import (
	device_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/manager/devices"
	"github.com/gin-gonic/gin"
)

type DeviceFactory struct {
	deviceController *device_controller.DeviceController
}

func NewDeviceFactory(controller *device_controller.DeviceController) *DeviceFactory {
	// TODO: apply validations for arguments if needed

	factory := &DeviceFactory{
		deviceController: controller,
	}

	return factory
}

func (f *DeviceFactory) Handle(engine *gin.RouterGroup) *gin.RouterGroup {
	deviceGroup := engine.Group("/devices")
	{
		deviceGroup.GET("/:device_id", f.deviceController.GetDevice)
		deviceGroup.POST("", f.deviceController.CreateDevice)
		deviceGroup.PUT("", f.deviceController.UpdateDevice)
		deviceGroup.DELETE("/:device_id", f.deviceController.DeleteDevice)
	}

	return deviceGroup
}
