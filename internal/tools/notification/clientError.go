package notification

import "github.com/gin-gonic/gin"

var (
	ClientError iError = &sError{}
)

type iError interface {
	Response(err error) gin.H
}

type sError struct {
	//
}

func (e *sError) Response(err error) gin.H {
	return gin.H{"error": err.Error()}
}
