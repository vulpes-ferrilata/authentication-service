package commands

type RevokeToken struct {
	RefreshToken string `validate:"required,jwt"`
}
