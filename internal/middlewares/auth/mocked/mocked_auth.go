package mocked_auth_middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	pkg_auth "github.com/Pietroski/SimpleTransactionDemo/pkg/tools/extractors/auth"
)

const (
	AuthorizationKey = "Authorization"
)

var (
	ErrToAssertVar = errors.New("error to type assert value from context")
)

func MockedAuthMiddleware(ctx *gin.Context) {
	rawBearToken := ctx.Request.Header.Get(AuthorizationKey)
	bearToken, err := pkg_auth.ExtractBearToken(rawBearToken)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
	}

	authValue, ok := MockedAuthMap[MockedBearToken(bearToken)]
	if !ok {
		_ = ctx.AbortWithError(http.StatusUnauthorized, pkg_auth.ErrInvalidAuthBearToken)
	}

	ctx.Set(CtxMockedAuthKey.String(), authValue)
	ctx.Next()
}

func MockedAuthMiddlewareExtractor(ctx *gin.Context) (MockedAuthValues, error) {
	rawAuthValue, ok := ctx.Get(CtxMockedAuthKey.String())
	if !ok {
		return MockedAuthValues{}, pkg_auth.ErrInvalidAuthBearToken
	}

	authValue, ok := rawAuthValue.(MockedAuthValues)
	if !ok {
		return MockedAuthValues{}, ErrToAssertVar
	}

	return authValue, nil
}
