package interceptors

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/context_values"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func NewErrorHandlerInterceptor(universalTranslator *ut.UniversalTranslator) *ErrorHandlerInterceptor {
	return &ErrorHandlerInterceptor{
		universalTranslator: universalTranslator,
	}
}

type ErrorHandlerInterceptor struct {
	universalTranslator *ut.UniversalTranslator
}

func (e ErrorHandlerInterceptor) ServerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		result, err := handler(ctx, req)
		if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			err = app_errors.NewValidationError(fieldErrors...)
		}
		if grpcErr, ok := errors.Cause(err).(app_errors.GrpcError); ok {
			locales := context_values.GetLocales(ctx)
			translator, _ := e.universalTranslator.FindTranslator(locales...)

			stt := grpcErr.Status(translator)

			return result, stt.Err()
		}
		if status, ok := status.FromError(errors.Cause(err)); ok {
			return result, status.Err()
		}
		if err != nil {
			return result, errors.WithStack(err)
		}

		return result, nil
	}
}
