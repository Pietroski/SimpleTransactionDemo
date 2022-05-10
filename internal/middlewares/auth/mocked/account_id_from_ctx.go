package mocked_auth_middleware

import (
	"net/http"

	"github.com/Pietroski/SimpleTransactionDemo/internal/tools/notification"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AccountIdCtxExtractor(ctx *gin.Context) (uuid.UUID, int, gin.H) {
	authInfo, err := MockedAuthMiddlewareExtractor(ctx)
	if err != nil {
		//if errors.As(err, &ErrToAssertVar) {
		//	return uuid.Nil,
		//		http.StatusInternalServerError,
		//		notification.ClientError.Response(err)
		//}

		// errors.As(err, &pkg_auth_extractorErrInvalidAuthBearerToken)
		return uuid.Nil,
			http.StatusBadRequest,
			notification.ClientError.Response(err)
	}

	rawAccountID := authInfo.AccountID
	accountID, err := uuid.Parse(rawAccountID.String())
	if err != nil {
		return uuid.Nil,
			http.StatusBadRequest,
			notification.ClientError.Response(err)
	}

	return accountID, 0, gin.H{}
}
