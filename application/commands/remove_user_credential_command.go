package commands

type RemoveUserCredentialCommand struct {
	ID string `validate:"required,uuid4"`
}
