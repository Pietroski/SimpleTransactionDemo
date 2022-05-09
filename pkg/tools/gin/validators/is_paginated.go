package pkg_gin_custom_validators

import (
	"github.com/gin-gonic/gin"
)

func IsPaginated(ctx gin.Context) bool {
	//ctx.Query()

	return true
}
