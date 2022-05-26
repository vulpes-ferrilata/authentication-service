package queries

type GetRefreshTokenByClaimQuery struct {
	UserID string `validate:"required,uuid4"`
	JTI    string `validate:"required,uuid4"`
}
