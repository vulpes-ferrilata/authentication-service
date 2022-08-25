package commands

type LoginCommand struct {
	ClaimID  string `validate:"required,objectid"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
