package commands

type DeleteUserCredential struct {
	UserCredentialID string `validate:"required,objectid"`
}
