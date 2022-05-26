package validators

import (
	"context"

	service "github.com/VulpesFerrilata/authentication-service"
	"github.com/VulpesFerrilata/authentication-service/persistence"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentialValidator interface {
	ValidateAuthentication(ctx context.Context, email string, password string) error
}

func NewUserCredentialValidator(userCredentialRepository repositories.UserCredentialRepository) UserCredentialValidator {
	return &userCredentialValidator{
		userCredentialRepository: userCredentialRepository,
	}
}

type userCredentialValidator struct {
	userCredentialRepository repositories.UserCredentialRepository
}

func (u userCredentialValidator) ValidateAuthentication(ctx context.Context, email string, password string) error {
	userCredentialEntity, err := u.userCredentialRepository.GetByEmail(ctx, email)
	if errors.Is(err, persistence.ErrRecordNotFound) {
		return errors.WithStack(service.ErrEmailIsInvalid)
	}
	if err != nil {
		return errors.WithStack(err)
	}

	if err := bcrypt.CompareHashAndPassword(userCredentialEntity.HashPassword, []byte(password)); err != nil {
		return errors.WithStack(service.ErrPasswordIsInvalid)
	}

	return nil
}
