package pkg_auth

import (
	"errors"
	"strings"
)

var (
	ErrInvalidAuthBearToken = errors.New("invalid authorization bear token")
)

func ExtractBearToken(rawBearToken string) (string, error) {
	tokenCompound := strings.Fields(rawBearToken)
	if len(tokenCompound) != 2 {
		return "", ErrInvalidAuthBearToken
	}

	return tokenCompound[1], nil
}
