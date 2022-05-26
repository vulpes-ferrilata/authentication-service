package repositories

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
)

type UserCredentialRepository interface {
	IsEmailExists(ctx context.Context, email string) (bool, error)
	GetByEmail(ctx context.Context, email string) (*entities.UserCredential, error)
	Repository[entities.UserCredential]
}
