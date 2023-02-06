package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/domain/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateUserCredentialCommand struct {
	UserCredentialID string `validate:"required,objectid"`
	UserID           string `validate:"required,objectid"`
	Email            string `validate:"required,email"`
	Password         string `validate:"required,gte=8"`
}

func NewCreateUserCredentialCommandHandler(userCredentialRepository repositories.UserCredentialRepository,
	userCredentialValidationService services.UserCredentialValidationService) *CreateUserCredentialCommandHandler {
	return &CreateUserCredentialCommandHandler{
		userCredentialRepository:        userCredentialRepository,
		userCredentialValidationService: userCredentialValidationService,
	}
}

type CreateUserCredentialCommandHandler struct {
	userCredentialRepository        repositories.UserCredentialRepository
	userCredentialValidationService services.UserCredentialValidationService
}

func (c CreateUserCredentialCommandHandler) Handle(ctx context.Context, createUserCredentialCommand *CreateUserCredentialCommand) error {
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
