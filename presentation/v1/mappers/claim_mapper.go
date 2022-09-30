package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
)

func ToClaimResponse(claim *models.Claim) *responses.Claim {
	if claim == nil {
		return nil
	}

	return &responses.Claim{
		ID:     claim.ID.Hex(),
		UserID: claim.UserID.Hex(),
	}
}
