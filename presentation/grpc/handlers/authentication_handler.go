package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/queries"
	"github.com/VulpesFerrilata/authentication-service/application/queries/dtos"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/bus"
	"github.com/VulpesFerrilata/shared/proto/authentication"
	"github.com/pkg/errors"
)

func NewAuthenticationHandler(queryBus bus.QueryBus) authentication.AuthenticationHandler {
	return &authenticationHandler{
		queryBus: queryBus,
	}
}

type authenticationHandler struct {
	queryBus bus.QueryBus
}

func (a authenticationHandler) GetClaimByAccessToken(ctx context.Context, getClaimByAccessTokenRequest *authentication.GetClaimByAccessTokenRequest, claimResponse *authentication.ClaimResponse) error {
	getClaimByAccessTokenQuery := &queries.GetClaimByAccessTokenQuery{
		AccessToken: getClaimByAccessTokenRequest.GetAccessToken(),
	}
	result, err := a.queryBus.Execute(ctx, getClaimByAccessTokenQuery)
	if err != nil {
		return errors.WithStack(err)
	}
	claimDTO := result.(*dtos.ClaimDTO)

	claimResponse.UserID = claimDTO.UserID

	return nil
}
