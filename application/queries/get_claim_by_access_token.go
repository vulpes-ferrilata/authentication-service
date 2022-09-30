package queries

type GetClaimByAccessToken struct {
	AccessToken string `validate:"required,jwt"`
}
