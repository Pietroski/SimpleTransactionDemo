package pkg_auth_extractor

import (
	"errors"
	"strings"
)

var (
	ErrInvalidAuthBearerToken = errors.New("invalid authorization Bearer token")
)

func ExtractBearerToken(rawBearerToken string) (string, error) {
	tokenCompound := strings.Fields(rawBearerToken)
	if len(tokenCompound) != 2 {
		return "", ErrInvalidAuthBearerToken
	}

	return tokenCompound[1], nil
}
