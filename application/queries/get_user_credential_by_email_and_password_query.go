package queries

type GetUserCredentialByEmailAndPasswordQuery struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
