package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/authentication"
)

func ToClaimResponse(claim *models.Claim) *authentication.ClaimResponse {
	if claim == nil {
		return nil
	}

	return &authentication.ClaimResponse{
		ID:     claim.ID.Hex(),
		UserID: claim.UserID.Hex(),
	}
}
