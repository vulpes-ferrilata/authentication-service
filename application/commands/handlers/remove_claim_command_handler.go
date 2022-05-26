package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/commands"
	domain_services "github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/domain/validators"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/services"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRemoveClaimCommandHandler(claimValidator validators.ClaimValidator,
	claimService domain_services.ClaimService,
	tokenFactoryService services.TokenFactoryService) results.CommandHandlerResult {
	commandHandler := &removeClaimCommandHandler{
		claimValidator:      claimValidator,
		claimService:        claimService,
		tokenFactoryService: tokenFactoryService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type removeClaimCommandHandler struct {
	claimValidator      validators.ClaimValidator
	claimService        domain_services.ClaimService
	tokenFactoryService services.TokenFactoryService
}

func (r removeClaimCommandHandler) GetCommand() interface{} {
	return &commands.RemoveClaimCommand{}
}

func (r removeClaimCommandHandler) Handle(ctx context.Context, command interface{}) error {
	removeClaimCommand := command.(*commands.RemoveClaimCommand)

	jti, err := uuid.Parse(removeClaimCommand.JTI)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := uuid.Parse(removeClaimCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.claimValidator.ValidateAuthentication(ctx, userID, jti); err != nil {
		return errors.WithStack(err)
	}

	claim, err := r.claimService.GetByUserID(ctx, userID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := r.claimService.Delete(ctx, claim); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
