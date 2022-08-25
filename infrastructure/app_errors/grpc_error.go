package app_errors

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc/status"
)

type GrpcError interface {
	Status(translator ut.Translator) *status.Status
}
