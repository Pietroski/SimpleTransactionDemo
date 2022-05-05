package notification

import "github.com/gin-gonic/gin"

var (
	ClientError iError = &clientError{}
)

type iError interface {
	Response(err error) gin.H
}

type clientError struct {
	//
}

func (e *clientError) Response(err error) gin.H {
	return gin.H{"error": err.Error()}
}
