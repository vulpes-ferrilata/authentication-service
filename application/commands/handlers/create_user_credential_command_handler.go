package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/domain/services"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command/wrappers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewCreateUserCredentialCommandHandler(validate *validator.Validate,
	db *mongo.Database,
	userCredentialRepository repositories.UserCredentialRepository,
	userCredentialValidationService services.UserCredentialValidationService) command.CommandHandler[*commands.CreateUserCredential] {
	handler := &createUserCredentialCommandHandler{
		userCredentialRepository:        userCredentialRepository,
		userCredentialValidationService: userCredentialValidationService,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.CreateUserCredential](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type createUserCredentialCommandHandler struct {
	userCredentialRepository        repositories.UserCredentialRepository
	userCredentialValidationService services.UserCredentialValidationService
}

func (c createUserCredentialCommandHandler) Handle(ctx context.Context, createUserCredentialCommand *commands.CreateUserCredential) error {
	userCredentialID, err := primitive.ObjectIDFromHex(createUserCredentialCommand.UserCredentialID)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := primitive.ObjectIDFromHex(createUserCredentialCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	isExists, err := c.userCredentialValidationService.IsEmailAlreadyExists(ctx, createUserCredentialCommand.Email)
	if err != nil {
		return errors.WithStack(err)
	}
	if isExists {
		return errors.WithStack(app_errors.ErrEmailIsAlreadyExists)
	}

	userCredential := models.UserCredentialBuilder{}.
		SetID(userCredentialID).
		SetUserID(userID).
		SetEmail(createUserCredentialCommand.Email).
		Create()

	if err := userCredential.SetPassword(createUserCredentialCommand.Password); err != nil {
		return errors.WithStack(err)
	}

	if err := c.userCredentialRepository.Insert(ctx, userCredential); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
