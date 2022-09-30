package queries

type GetTokenByClaimID struct {
	ClaimID string `validate:"required,objectid"`
}
