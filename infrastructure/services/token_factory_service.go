package services

type TokenType string

var (
	AccessToken  TokenType = "access token"
	RefreshToken TokenType = "refresh token"
)

type TokenFactoryService interface {
	GetTokenService(tokenType TokenType) TokenService
}

func NewTokenFactoryService(accessTokenService TokenService,
	refreshTokenService TokenService) TokenFactoryService {
	return &tokenFactoryService{
		accessTokenService:  accessTokenService,
		refreshTokenService: refreshTokenService,
	}
}

type tokenFactoryService struct {
	accessTokenService  TokenService
	refreshTokenService TokenService
}

func (t tokenFactoryService) GetTokenService(tokenType TokenType) TokenService {
	switch tokenType {
	case AccessToken:
		return t.accessTokenService
	case RefreshToken:
		return t.refreshTokenService
	}

	return nil
}
