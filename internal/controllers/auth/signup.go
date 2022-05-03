package auth_controller

import (
	auth_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/auth"
	"github.com/gin-gonic/gin"
)

type signUpController struct {
	store auth_store.Store
}

func newSignUpController(store auth_store.Store) *signUpController {
	// TODO: apply validations for arguments if needed

	controller := &signUpController{
		store: store,
	}

	return controller
}

func (c *signUpController) CreateUser(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) GetUser(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) getUserByID(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) getUserByEmail(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) ListUsers(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) listUsers(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) listPaginatedUsers(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) DeleteUser(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) deleteUserByID(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) deleteUserByEmail(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) UpdateUser(ctx *gin.Context) {
	// TODO: implement me!!
}

func (c *signUpController) updateUser(ctx *gin.Context) {
	// TODO: implement me!!
}
