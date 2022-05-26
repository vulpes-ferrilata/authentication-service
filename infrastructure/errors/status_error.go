package errors

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewStatusError(stt *status.Status) AppError {
	return &statusError{
		stt: stt,
	}
}

type statusError struct {
	stt *status.Status
}

func (s statusError) Error() string {
	return s.stt.Err().Error()
}

func (s statusError) Problem(translator ut.Translator) (iris.Problem, error) {
	problem := iris.NewProblem()

	switch s.stt.Code() {
	case codes.InvalidArgument:
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
	case codes.FailedPrecondition:
		problem.Status(iris.StatusUnprocessableEntity)

		title, err := translator.T("business-error-title")
		if err != nil {
			return nil, errors.Wrap(err, "business-error-title")
		}
		problem.Title(title)

		detail, err := translator.T("business-error-detail")
		if err != nil {
			return nil, errors.Wrap(err, "business-error-detail")
		}
		problem.Detail(detail)
	default:
		problem.Status(iris.StatusInternalServerError)
		problem.Title("something went wrong")
		problem.Detail(s.stt.Message())
	}

	for _, detail := range s.stt.Details() {
		if badRequest, ok := detail.(*errdetails.BadRequest); ok {
			messages := make(map[string][]string)
			for _, fieldViolation := range badRequest.GetFieldViolations() {
				messages[fieldViolation.GetField()] = append(messages[fieldViolation.GetField()], fieldViolation.GetDescription())
			}
			problem.Key("errors", messages)
		}

		if preconditionFailure, ok := detail.(*errdetails.PreconditionFailure); ok {
			messages := make([]string, 0)
			for _, violation := range preconditionFailure.GetViolations() {
				message := violation.GetDescription()
				messages = append(messages, message)
			}
			problem.Key("errors", messages)
		}
	}

	return problem, nil
}

func (s statusError) Status(translator ut.Translator, serviceName string) (*status.Status, error) {
	return s.stt, nil
}
