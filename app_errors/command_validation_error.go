package app_errors

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewCommandValidationError(validationErrors validator.ValidationErrors) AppError {
	return &commandValidationError{
		validationErrors: validationErrors,
	}
}

type commandValidationError struct {
	validationErrors validator.ValidationErrors
}

func (c commandValidationError) Error() string {
	builder := new(strings.Builder)

	builder.WriteString("the command contains invalid parameters")

	for _, fieldError := range c.validationErrors {
		builder.WriteString("\n")
		builder.WriteString(fieldError.Error())
	}

	return builder.String()
}

func (c commandValidationError) Status(translator ut.Translator) *status.Status {
	detail, err := translator.T("command-validation-error")
	if err != nil {
		detail = "command-validation-error"
	}

	stt := status.New(codes.InvalidArgument, detail)

	badRequest := &errdetails.BadRequest{}
	for _, fieldError := range c.validationErrors {
		fieldViolation := &errdetails.BadRequest_FieldViolation{
			Field:       fieldError.Field(),
			Description: fieldError.Translate(translator),
		}
		badRequest.FieldViolations = append(badRequest.FieldViolations, fieldViolation)
	}

	stt, _ = stt.WithDetails(badRequest)

	return stt
}
