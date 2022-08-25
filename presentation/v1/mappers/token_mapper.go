package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/authentication"
)

func ToTokenResponse(token *models.Token) *authentication.TokenResponse {
	if token == nil {
		return nil
	}

	return &authentication.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}
