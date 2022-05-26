package service

import (
	app_errors "github.com/VulpesFerrilata/authentication-service/infrastructure/errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
)

var (
	ErrEmailIsInvalid                        = NewDetailError("invalid-email")
	ErrPasswordIsInvalid                     = NewDetailError("invalid-password")
	ErrTokenHasBeenExpiredOrRevoked          = NewDetailError("revoked-token")
	ErrAccountHasBeenLoggedInByAnotherDevice = NewDetailError("duplicate-login")
	ErrTokenIsInvalid                        = NewDetailError("invalid-token")
)

func NewDetailError(translationKey string) app_errors.DetailError {
	return &detailError{
		translationKey: translationKey,
	}
}

type detailError struct {
	translationKey string
}

func (d detailError) Error() string {
	return d.translationKey
}

func (d detailError) Translate(translator ut.Translator) (string, error) {
	message, err := translator.T(d.translationKey)
	if err != nil {
		return "", errors.Wrap(err, d.translationKey)
	}

	return message, nil
}
