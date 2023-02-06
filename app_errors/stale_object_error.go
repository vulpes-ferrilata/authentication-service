package app_errors

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewStaleObjectError(object string) AppError {
	return &staleObjectError{
		object: object,
	}
}

type staleObjectError struct {
	object string
}

func (s staleObjectError) Error() string {
	return fmt.Sprintf("attempted to update or delete stale %s", s.object)
}

func (s staleObjectError) Status(translator ut.Translator) *status.Status {
	object, err := translator.T(s.object)
	if err != nil {
		object = s.object
	}

	detail, err := translator.T("stale-object-error", object)
	if err != nil {
		detail = fmt.Sprintf("%s: %s", "stale-object-error", object)
	}

	stt := status.New(codes.Aborted, detail)

	return stt
}
