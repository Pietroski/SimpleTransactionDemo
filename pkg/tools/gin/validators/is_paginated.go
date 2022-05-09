package pkg_gin_custom_validators

import (
	"github.com/gin-gonic/gin"
)

type (
	Pagination struct {
		//
	}

	PaginatedQuery struct {
		//
	}
)

func IsCorrectlyPaginated(ctx gin.Context) (*Pagination, bool) {
	ctx.ShouldBindQuery()

	return true
}
