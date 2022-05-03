package auth_factory

import (
	auth_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/auth"
	"github.com/gin-gonic/gin"
)

type SignInFactory struct {
	authController *auth_controller.AuthController
}

func newSignInFactory(controller *auth_controller.AuthController) *SignInFactory {
	// TODO: apply validations for arguments if needed

	factory := &SignInFactory{
		authController: controller,
	}

	return factory
}

func (f *SignInFactory) Handle(engine *gin.RouterGroup) *gin.RouterGroup {
	signInGroup := engine.Group("/sign-in")
	{
		signInGroup.GET("", nil)
		signInGroup.POST("", nil)
		signInGroup.PUT("", nil)
		signInGroup.PATCH("", nil)
		signInGroup.DELETE("", nil)
	}

	return signInGroup
}
