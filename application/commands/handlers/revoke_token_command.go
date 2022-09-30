package handlers

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command/wrappers"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRevokeTokenCommandHandler(validate *validator.Validate,
	db *mongo.Database,
	tokenServiceResolver services.TokenServiceResolver,
	claimRepository repositories.ClaimRepository) command.CommandHandler[*commands.RevokeToken] {
	handler := &revokeTokenCommandHandler{
		tokenServiceResolver: tokenServiceResolver,
		claimRepository:      claimRepository,
	}
	transactionWrapper := wrappers.NewTransactionWrapper[*commands.RevokeToken](db, handler)
	validationWrapper := wrappers.NewValidationWrapper(validate, transactionWrapper)

	return validationWrapper
}

type revokeTokenCommandHandler struct {
	tokenServiceResolver services.TokenServiceResolver
	claimRepository      repositories.ClaimRepository
}

func (r revokeTokenCommandHandler) Handle(ctx context.Context, revokeTokenCommand *commands.RevokeToken) error {
	id, err := r.tokenServiceResolver.GetTokenService(services.RefreshToken).Decrypt(ctx, revokeTokenCommand.RefreshToken)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.claimRepository.Delete(ctx, id); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
