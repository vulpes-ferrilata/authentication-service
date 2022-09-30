package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginCommandHandler(validate *validator.Validate,
	db *mongo.Database,
	userCredentialRepository repositories.UserCredentialRepository,
	claimRepository repositories.ClaimRepository) command.CommandHandler[*commands.Login] {
	handler := &loginCommandHandler{
		userCredentialRepository: userCredentialRepository,
		claimRepository:          claimRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.Login](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper

}

type loginCommandHandler struct {
	userCredentialRepository repositories.UserCredentialRepository
	claimRepository          repositories.ClaimRepository
}

func (l loginCommandHandler) Handle(ctx context.Context, loginCommand *commands.Login) error {
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
