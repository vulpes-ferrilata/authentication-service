package services

import "context"

type UserCredentialValidationService interface {
	IsEmailAlreadyExists(ctx context.Context, email string) (bool, error)
}
