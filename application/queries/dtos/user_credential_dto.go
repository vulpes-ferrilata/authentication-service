package dtos

import (
	"github.com/VulpesFerrilata/authentication-service/domain/models"
)

func NewUserCredentialDTO(userCredential *models.UserCredential) *UserCredentialDTO {
	return &UserCredentialDTO{
		ID:     userCredential.GetID().String(),
		UserID: userCredential.GetUserID().String(),
	}
}

type UserCredentialDTO struct {
	ID     string
	UserID string
}
