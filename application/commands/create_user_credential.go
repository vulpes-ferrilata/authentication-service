package commands

type CreateUserCredential struct {
	UserCredentialID string `validate:"required,objectid"`
	UserID           string `validate:"required,objectid"`
	Email            string `validate:"required,email"`
	Password         string `validate:"required,gte=8"`
}
