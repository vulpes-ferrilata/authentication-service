package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
)

type ClaimMapper struct{}

func (c ClaimMapper) ToResponse(claim *models.Claim) (*responses.Claim, error) {
	if claim == nil {
		return nil, nil
	}

	return &responses.Claim{
		ID:     claim.ID.Hex(),
		UserID: claim.UserID.Hex(),
	}, nil
}
