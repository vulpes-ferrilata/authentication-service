package queries

type GetClaimByAccessTokenQuery struct {
	AccessToken string `validate:"required,jwt"`
}
