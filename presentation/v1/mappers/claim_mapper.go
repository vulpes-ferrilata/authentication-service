package mappers

import (
	pb_models "github.com/vulpes-ferrilata/authentication-service-proto/pb/models"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
)

type ClaimMapper struct{}

func (c ClaimMapper) ToResponse(claim *models.Claim) (*pb_models.Claim, error) {
	if claim == nil {
		return nil, nil
	}

	return &pb_models.Claim{
		ID:     claim.ID.Hex(),
		UserID: claim.UserID.Hex(),
	}, nil
}
