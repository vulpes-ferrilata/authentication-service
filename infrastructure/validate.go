package infrastructure

import (
	"context"
	"strings"

	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
)

var (
	ErrTranslatorNotFound = errors.New("translator not found")
)

func NewValidate(universalTranslator *ut.UniversalTranslator, userCredentialRepository repositories.UserCredentialRepository) (*validator.Validate, error) {
	validate := validator.New()

	en := en.New()

	translator, ok := universalTranslator.GetTranslator(en.Locale())
	if !ok {
		return nil, errors.Wrap(ErrTranslatorNotFound, en.Locale())
	}

	if err := en_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		return nil, errors.WithStack(err)
	}

	if err := validate.RegisterValidationCtx("duplicate", func(ctx context.Context, fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ".")
		if len(params) != 2 {
			return false
		}

		table := params[0]
		column := params[1]

		if table == "user_credentials" && column == "email" {
			isExists, err := userCredentialRepository.IsEmailExists(ctx, fl.Field().String())
			if err != nil {
				return false
			}
			if isExists {
				return false
			}

			return true
		}

		return false
	}); err != nil {
		return nil, errors.WithStack(err)
	}

	return validate, nil
}
