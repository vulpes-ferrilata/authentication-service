package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service/domain/models"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/domain/mongo/documents"
)

type UserCredentialMapper struct{}

func (u UserCredentialMapper) ToDocument(userCredential *models.UserCredential) (*documents.UserCredential, error) {
	if userCredential == nil {
		return nil, nil
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
	}, nil
}

func (u UserCredentialMapper) ToDomain(userCredentialDocument *documents.UserCredential) (*models.UserCredential, error) {
	if userCredentialDocument == nil {
		return nil, nil
	}

	return models.UserCredentialBuilder{}.
		SetID(userCredentialDocument.ID).
		SetUserID(userCredentialDocument.UserID).
		SetEmail(userCredentialDocument.Email).
		SetHashPassword(userCredentialDocument.HashPassword).
		SetVersion(userCredentialDocument.Version).
		Create(), nil
}
