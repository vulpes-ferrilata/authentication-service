package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/documents"
)

func ToUserCredentialDocument(userCredential *models.UserCredential) *documents.UserCredential {
	if userCredential == nil {
		return nil
	}

	return &documents.UserCredential{
		DocumentRoot: documents.DocumentRoot{
			Document: documents.Document{
				ID: userCredential.GetID(),
			},
		},
		UserID:       userCredential.GetUserID(),
		Email:        userCredential.GetEmail(),
		HashPassword: userCredential.GetHashPassword(),
		Version:      userCredential.GetVersion(),
	}
}

func ToUserCredentialDomain(userCredentialDocument *documents.UserCredential) *models.UserCredential {
	if userCredentialDocument == nil {
		return nil
	}

	return models.NewUserCredentialBuilder().
		SetID(userCredentialDocument.ID).
		SetUserID(userCredentialDocument.UserID).
		SetEmail(userCredentialDocument.Email).
		SetHashPassword(userCredentialDocument.HashPassword).
		SetVersion(userCredentialDocument.Version).
		Create()
}
