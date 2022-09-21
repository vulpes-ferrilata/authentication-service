package commands

type RevokeTokenCommand struct {
	RefreshToken string `validate:"required,jwt"`
}
