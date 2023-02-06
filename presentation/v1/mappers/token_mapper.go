package mappers

import (
	pb_models "github.com/vulpes-ferrilata/authentication-service-proto/pb/models"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
)

type TokenMapper struct{}

func (t TokenMapper) ToResponse(token *models.Token) (*pb_models.Token, error) {
	if token == nil {
		return nil, nil
	}

	return &pb_models.Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}, nil
}
