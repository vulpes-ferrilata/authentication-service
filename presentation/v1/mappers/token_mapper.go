package mappers

import (
	"github.com/vulpes-ferrilata/authentication-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
)

type TokenMapper struct{}

func (t TokenMapper) ToResponse(token *models.Token) (*responses.Token, error) {
	if token == nil {
		return nil, nil
	}

	return &responses.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
