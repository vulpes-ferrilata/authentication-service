package queries

type GetTokenByRefreshTokenQuery struct {
	RefreshToken string `validate:"required,jwt"`
}
