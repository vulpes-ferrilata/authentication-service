package queries

type GetClaimByRefreshTokenQuery struct {
	RefreshToken string `validate:"required,jwt"`
}
