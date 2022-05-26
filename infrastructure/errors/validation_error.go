package errors

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewValidationError(fieldErrors ...validator.FieldError) AppError {
	return &validationError{
		fieldErrors: fieldErrors,
	}
}

type validationError struct {
	fieldErrors []validator.FieldError
}

func (v validationError) Error() string {
	builder := new(strings.Builder)

	builder.WriteString("one or more validation errors occured")

	for _, fieldError := range v.fieldErrors {
		builder.WriteString("\n")
		builder.WriteString(fieldError.Error())
	}

	return builder.String()
}

func (v validationError) Problem(translator ut.Translator) (iris.Problem, error) {
	problem := iris.NewProblem()
	problem.Status(iris.StatusBadRequest)

	title, err := translator.T("validation-error-title")
	if err != nil {
		return nil, errors.Wrap(err, "validation-error-title")
	}
	problem.Title(title)

	detail, err := translator.T("validation-error-detail")
	if err != nil {
		return nil, errors.Wrap(err, "validation-error-detail")
	}
	problem.Detail(detail)

	messages := make(map[string][]string)
	for _, fieldError := range v.fieldErrors {
		message := fieldError.Translate(translator)
		messages[fieldError.Namespace()] = append(messages[fieldError.Namespace()], message)
	}
	problem.Key("errors", messages)

	return problem, nil
}

func (v validationError) Status(translator ut.Translator, serviceName string) (*status.Status, error) {
	stt := status.New(codes.InvalidArgument, serviceName)

	badRequest := &errdetails.BadRequest{}
	for _, fieldError := range v.fieldErrors {
		fieldViolation := &errdetails.BadRequest_FieldViolation{
			Field:       fieldError.Namespace(),
			Description: fieldError.Translate(translator),
		}
		badRequest.FieldViolations = append(badRequest.FieldViolations, fieldViolation)
	}
	stt, err := stt.WithDetails(badRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stt, nil

}
