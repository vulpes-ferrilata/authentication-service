package commands

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/domain/repositories"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/services"
)

type RevokeTokenCommand struct {
	RefreshToken string `validate:"required,jwt"`
}

func NewRevokeTokenCommandHandler(tokenServiceResolver services.TokenServiceResolver,
	claimRepository repositories.ClaimRepository) *RevokeTokenCommandHandler {
	return &RevokeTokenCommandHandler{
		tokenServiceResolver: tokenServiceResolver,
		claimRepository:      claimRepository,
	}
}

type RevokeTokenCommandHandler struct {
	tokenServiceResolver services.TokenServiceResolver
	claimRepository      repositories.ClaimRepository
}

func (r RevokeTokenCommandHandler) Handle(ctx context.Context, revokeTokenCommand *RevokeTokenCommand) error {
	id, err := r.tokenServiceResolver.GetTokenService(services.RefreshToken).Decrypt(ctx, revokeTokenCommand.RefreshToken)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.claimRepository.Delete(ctx, id); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
