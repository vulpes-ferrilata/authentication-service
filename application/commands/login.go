package commands

type Login struct {
	ClaimID  string `validate:"required,objectid"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
