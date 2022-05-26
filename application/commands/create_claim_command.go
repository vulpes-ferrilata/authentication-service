package commands

type CreateClaimCommand struct {
	JTI    string `validate:"required,uuid4"`
	UserID string `validate:"required,uuid4"`
}
