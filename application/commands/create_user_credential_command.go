package commands

type CreateUserCredentialCommand struct {
	ID       string `validate:"required,objectid"`
	UserID   string `validate:"required,objectid"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,gte=8"`
}
