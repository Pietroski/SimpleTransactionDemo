package auth_controller

import (
	auth_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/auth"
)

type AuthController struct {
	SignUp   *signUpController
	SignIn   *signInController
	Recovery *recoveryController
}

func NewAuthController(store auth_store.Store) *AuthController {
	// TODO: apply validations for arguments if needed

	controller := &AuthController{
		SignUp:   newSignUpController(store),
		SignIn:   newSignInController(),
		Recovery: newRecoveryController(),
	}

	return controller
}
