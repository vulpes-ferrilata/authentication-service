package app_errors

var (
	ErrUserCredentialNotFound  = NewNotFoundError("user-credential")
	ErrStaleUserCredential     = NewStaleObjectError("user-credential")
	ErrClaimNotFound           = NewNotFoundError("claim")
	ErrUnableToEncryptPassword = NewBusinessRuleError("unable-to-encrypt-password")
	ErrPasswordIsInvalid       = NewBusinessRuleError("password-is-invalid")
	ErrTokenHasBeenExpired     = NewAuthenticationError("token-has-been-expired")
	ErrTokenHasBeenRevoked     = NewAuthenticationError("token-has-been-revoked")
	ErrTokenIsInvalid          = NewAuthenticationError("token-is-invalid")
)
