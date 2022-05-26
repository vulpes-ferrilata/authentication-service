package repositories

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
	"github.com/google/uuid"
)

type Repository[T entities.Entities] interface {
	GetByID(ctx context.Context, id uuid.UUID) (*T, error)
	Save(ctx context.Context, entity *T) error
	Delete(ctx context.Context, entity *T) error
}
