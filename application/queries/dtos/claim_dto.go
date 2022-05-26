package dtos

import (
	"github.com/VulpesFerrilata/authentication-service/domain/models"
)

func NewClaimDTO(claim *models.Claim) *ClaimDTO {
	return &ClaimDTO{
		UserID: claim.GetUserID().String(),
		JTI:    claim.GetJTI().String(),
	}
}

type ClaimDTO struct {
	UserID string
	JTI    string
}
