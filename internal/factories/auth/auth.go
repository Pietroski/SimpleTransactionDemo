package auth_factory

import (
	auth_store "github.com/Pietroski/SimpleTransactionDemo/internal/adaptors/datastore/postgresql/sqlc/auth"
	auth_controller "github.com/Pietroski/SimpleTransactionDemo/internal/controllers/auth"
	"github.com/gin-gonic/gin"
)

type AuthServer struct {
	address   string
	router    *gin.Engine
	authStore auth_store.Store
}

func NewAuthServer(store auth_store.Store) *AuthServer {
	// TODO: apply validations for arguments if needed

	factory := &AuthServer{
		authStore: store,
	}

	factory.Handle()

	return factory
}

func (f *AuthServer) Handle() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// TODO: implement error handling for CORS
	_ = r.SetTrustedProxies([]string{"localhost", "127.0.0.1"})

	authController := auth_controller.NewAuthController(f.authStore)
	signUpFactory := newSignUpFactory(authController)
	signInFactory := newSignInFactory(authController)
	recoveryFactory := newRecoveryFactory(authController)

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			signUpFactory.Handle(auth)
			signInFactory.Handle(auth)
			recoveryFactory.Handle(auth)
		}
	}

	f.router = r
}

func (f *AuthServer) Start() error {
	f.address = ":8088"
	return f.router.Run(f.address)
}
