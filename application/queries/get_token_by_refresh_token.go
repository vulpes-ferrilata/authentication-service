package queries

type GetTokenByRefreshToken struct {
	RefreshToken string `validate:"required,jwt"`
}
