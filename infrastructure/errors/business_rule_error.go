package errors

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBusinessRuleError(detailErrors ...DetailError) AppError {
	return &businessRuleError{
		detailErrors: detailErrors,
	}
}

type businessRuleError struct {
	detailErrors []DetailError
}

func (b businessRuleError) Error() string {
	builder := new(strings.Builder)

	builder.WriteString("business rule violation occured")

	for _, detailError := range b.detailErrors {
		builder.WriteString("\n")
		builder.WriteString(detailError.Error())
	}

	return builder.String()
}

func (b businessRuleError) Problem(translator ut.Translator) (iris.Problem, error) {
	problem := iris.NewProblem()
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

	messages := make([]string, 0)
	for _, detailError := range b.detailErrors {
		message, err := detailError.Translate(translator)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		messages = append(messages, message)
	}
	problem.Key("errors", messages)

	return problem, nil
}

func (b businessRuleError) Status(translator ut.Translator, serviceName string) (*status.Status, error) {
	stt := status.New(codes.FailedPrecondition, serviceName)

	preconditionFailure := &errdetails.PreconditionFailure{}
	for _, detailError := range b.detailErrors {
		message, err := detailError.Translate(translator)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		violation := &errdetails.PreconditionFailure_Violation{
			Description: message,
		}

		preconditionFailure.Violations = append(preconditionFailure.Violations, violation)
	}
	stt, err := stt.WithDetails(preconditionFailure)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return stt, nil
}
