package mocked_validator_v10

import "github.com/go-playground/validator/v10"

type FieldLevelValidator interface {
	validator.FieldLevel
}
