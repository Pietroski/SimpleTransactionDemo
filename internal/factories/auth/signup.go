package auth_factory

import (
	auth_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/auth"
	"github.com/gin-gonic/gin"
)

type SignUpFactory struct {
	authController *auth_controller.AuthController
}

func newSignUpFactory(controller *auth_controller.AuthController) *SignUpFactory {
	// TODO: apply validations for arguments if needed

	factory := &SignUpFactory{
		authController: controller,
	}

	return factory
}

func (f *SignUpFactory) Handle(engine *gin.RouterGroup) *gin.RouterGroup {
	signUpGroup := engine.Group("/sign-up")
	{
		signUpGroup.GET("/:user_id", f.authController.SignUp.GetUser)
		signUpGroup.POST("", f.authController.SignUp.CreateUser)
		signUpGroup.PUT("", f.authController.SignUp.UpdateUser)
		signUpGroup.DELETE("/:user_id", f.authController.SignUp.DeleteUser)
	}

	return signUpGroup
}
