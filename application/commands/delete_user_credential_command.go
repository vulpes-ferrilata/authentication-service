package commands

type DeleteUserCredentialCommand struct {
	ID string `validate:"required,objectid"`
}
