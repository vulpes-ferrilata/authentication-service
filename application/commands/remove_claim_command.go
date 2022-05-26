package commands

type RemoveClaimCommand struct {
	JTI    string `validate:"required,uuid4"`
	UserID string `validate:"required,uuid4"`
}
