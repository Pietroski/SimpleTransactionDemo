package server

type (
	Server interface {
		Handle()
		Start() error
	}
)
