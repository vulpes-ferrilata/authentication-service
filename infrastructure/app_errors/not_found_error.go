package app_errors

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewNotFoundError(object string) AppError {
	return &notFoundError{
		object: object,
	}
}

type notFoundError struct {
	object string
}

func (n notFoundError) Error() string {
	return fmt.Sprintf("%s not found", n.object)
}

func (n notFoundError) Status(translator ut.Translator) *status.Status {
	object, err := translator.T(n.object)
	if err != nil {
		object = n.object
	}

	detail, err := translator.T("not-found-error", object)
	if err != nil {
		detail = n.Error()
	}

	stt := status.New(codes.NotFound, detail)

	return stt
}
