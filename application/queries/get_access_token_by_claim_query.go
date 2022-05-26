package queries

type GetAccessTokenByClaimQuery struct {
	UserID string `validate:"required,uuid4"`
	JTI    string `validate:"required,uuid4"`
}
