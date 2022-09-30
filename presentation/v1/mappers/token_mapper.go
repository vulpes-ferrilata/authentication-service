package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
)

func ToTokenResponse(token *models.Token) *responses.Token {
	if token == nil {
		return nil
	}

	return &responses.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}
