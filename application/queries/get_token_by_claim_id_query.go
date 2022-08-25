package queries

type GetTokenByClaimIDQuery struct {
	ClaimID string `validate:"required,objectid"`
}
