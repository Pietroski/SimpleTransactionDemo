package auth_factory

import (
	auth_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/auth"
	"github.com/gin-gonic/gin"
)

type RecoveryFactory struct {
	authController *auth_controller.AuthController
}

func newRecoveryFactory(controller *auth_controller.AuthController) *RecoveryFactory {
	// TODO: apply validations for arguments if needed

	factory := &RecoveryFactory{
		authController: controller,
	}

	return factory
}

func (f *RecoveryFactory) Handle(engine *gin.RouterGroup) *gin.RouterGroup {
	signUpGroup := engine.Group("/recovery")
	{
		signUpGroup.GET("/:user_id", nil)
		signUpGroup.POST("", nil)
	}

	return signUpGroup
}
