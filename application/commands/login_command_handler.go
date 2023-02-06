package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginCommand struct {
	ClaimID  string `validate:"required,objectid"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func NewLoginCommandHandler(userCredentialRepository repositories.UserCredentialRepository,
	claimRepository repositories.ClaimRepository) *LoginCommandHandler {
	return &LoginCommandHandler{
		userCredentialRepository: userCredentialRepository,
		claimRepository:          claimRepository,
	}

}

type LoginCommandHandler struct {
	userCredentialRepository repositories.UserCredentialRepository
	claimRepository          repositories.ClaimRepository
}

func (l LoginCommandHandler) Handle(ctx context.Context, loginCommand *LoginCommand) error {
	userCredential, err := l.userCredentialRepository.GetByEmail(ctx, loginCommand.Email)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := userCredential.ComparePassword(loginCommand.Password); err != nil {
		return errors.WithStack(err)
	}

	id, err := primitive.ObjectIDFromHex(loginCommand.ClaimID)
	if err != nil {
		return errors.WithStack(err)
	}

	claim := models.ClaimBuilder{}.
		SetID(id).
		SetUserID(userCredential.GetUserID()).
		Create()

	if err := l.claimRepository.Insert(ctx, claim); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
