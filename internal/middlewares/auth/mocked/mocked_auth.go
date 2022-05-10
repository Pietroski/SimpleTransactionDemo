package mocked_auth_middleware

import (
	"errors"
	"net/http"

	pkg_auth_extractor "github.com/Pietroski/SimpleTransactionDemo/pkg/tools/extractors/auth"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationKey = "Authorization"
)

var (
	ErrToAssertVar = errors.New("error to type assert value from context")
)

func MockedAuthMiddleware(ctx *gin.Context) {
	rawBearerToken := ctx.Request.Header.Get(AuthorizationKey)
	BearerToken, err := pkg_auth_extractor.ExtractBearerToken(rawBearerToken)
	if err != nil {
		_ = ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	authValue, ok := MockedAuthMap[MockedBearerToken(BearerToken)]
	if !ok {
		_ = ctx.AbortWithError(http.StatusUnauthorized, pkg_auth_extractor.ErrInvalidAuthBearerToken)
		return
	}

	ctx.Set(CtxMockedAuthKey.String(), authValue)
	ctx.Next()
}

func MockedAuthMiddlewareExtractor(ctx *gin.Context) (MockedAuthValues, error) {
	rawAuthValue, ok := ctx.Get(CtxMockedAuthKey.String())
	if !ok {
		return MockedAuthValues{}, pkg_auth_extractor.ErrInvalidAuthBearerToken
	}

	authValue, ok := rawAuthValue.(MockedAuthValues)
	if !ok {
		return MockedAuthValues{}, ErrToAssertVar
	}

	return authValue, nil
}
