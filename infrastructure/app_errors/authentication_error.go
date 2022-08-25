package app_errors

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAuthenticationError(message string, params ...string) AppError {
	return &authenticationError{
		message: message,
		params:  params,
	}
}

type authenticationError struct {
	message string
	params  []string
}

func (a authenticationError) Error() string {
	return fmt.Sprintf("authentication failed: %s", a.message)
}

func (a authenticationError) Status(translator ut.Translator) *status.Status {
	detail, err := translator.T(a.message, a.params...)
	if err != nil {
		detail = a.message
	}

	stt := status.New(codes.Unauthenticated, detail)

	return stt
}
