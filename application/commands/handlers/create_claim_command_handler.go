package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/commands"
	domain_services "github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewCreateClaimCommandHandler(claimService domain_services.ClaimService) results.CommandHandlerResult {
	commandHandler := &createClaimCommandHandler{
		claimService: claimService,
	}

	return results.CommandHandlerResult{
		CommandHandler: commandHandler,
	}
}

type createClaimCommandHandler struct {
	claimService domain_services.ClaimService
}

func (c createClaimCommandHandler) GetCommand() interface{} {
	return &commands.CreateClaimCommand{}
}

func (c createClaimCommandHandler) Handle(ctx context.Context, command interface{}) error {
	createClaimCommand := command.(*commands.CreateClaimCommand)

	jti, err := uuid.Parse(createClaimCommand.JTI)
	if err != nil {
		return errors.WithStack(err)
	}

	userID, err := uuid.Parse(createClaimCommand.UserID)
	if err != nil {
		return errors.WithStack(err)
	}

	claim, err := c.claimService.NewClaim(ctx, jti, userID)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := c.claimService.Save(ctx, claim); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
