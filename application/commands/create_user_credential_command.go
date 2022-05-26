package commands

type CreateUserCredentialCommand struct {
	ID       string `validate:"required,uuid4"`
	UserID   string `validate:"required,uuid4"`
	Email    string `validate:"required,email,duplicate=user_credentials.email"`
	Password string `validate:"required,gte=8"`
}
