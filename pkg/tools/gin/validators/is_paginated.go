package pkg_gin_custom_validators

import (
	"errors"
	"net/http"

	"github.com/Pietroski/SimpleTransactionDemo/internal/tools/notification"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalid = errors.New("invalid pagination parameters")
)

type (
	Pagination struct {
		Limit  int32
		Offset int32
	}

	QueryPagination struct {
		PageID   int32 `form:"page_id"`
		PageSize int32 `form:"page_size"`
	}
)

func IsCorrectlyPaginated(ctx *gin.Context) (p *Pagination, statusCode int, ginResp gin.H) {
	var qp QueryPagination
	if err := ctx.ShouldBindQuery(&qp); err != nil {
		return nil,
			http.StatusInternalServerError,
			notification.ClientError.Response(err)
	}

	if qp.PageSize == 0 && qp.PageID == 0 ||
		qp.PageSize != 0 && qp.PageID != 0 {
		return &Pagination{
			Limit:  qp.PageSize,
			Offset: (qp.PageID - 1) * qp.PageSize,
		}, 0, nil
	}

	return nil,
		http.StatusBadRequest,
		notification.ClientError.Response(ErrInvalid)
}

func IsPaginated(p *Pagination) bool {
	return !(p.Limit == 0 && p.Offset == 0)
}
