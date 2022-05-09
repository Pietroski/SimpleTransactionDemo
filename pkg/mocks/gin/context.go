package mocked_gin

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type Context struct {
	gin.Context
}

func (c *Context) JSON(statusCode int, obj interface{}) {
	bb, err := json.MarshalIndent(obj, "", " ")
	log.Println(statusCode, "=>", string(bb), err)
}
