package repositories

import (
	"context"

	infrastructure_context "github.com/VulpesFerrilata/authentication-service/infrastructure/context"
	"github.com/VulpesFerrilata/authentication-service/persistence"
	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewRepository[T entities.Entities]() repositories.Repository[T] {
	return &repository[T]{}
}

type repository[T entities.Entities] struct{}

func (r repository[T]) GetByID(ctx context.Context, id uuid.UUID) (*T, error) {
	entity := new(T)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.First(entity, id)
	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(persistence.ErrRecordNotFound)
	}
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return entity, nil
}

func (r repository[T]) IsExists(ctx context.Context, entity *T) (bool, error) {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(entity).First(nil)
	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return true, nil
}

func (r repository[T]) Insert(ctx context.Context, entity *T) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Create(entity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (r repository[T]) Update(ctx context.Context, entity *T) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Updates(entity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}

func (r repository[T]) Save(ctx context.Context, entity *T) error {
	isExists, err := r.IsExists(ctx, entity)
	if err != nil {
		return errors.WithStack(err)
	}

	if isExists {
		if err := r.Update(ctx, entity); err != nil {
			return errors.WithStack(err)
		}
	} else {
		if err := r.Insert(ctx, entity); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (g repository[T]) Delete(ctx context.Context, entity *T) error {
	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	tx = tx.Delete(entity)
	if err := tx.Error; err != nil {
		return errors.WithStack(err)
	}
	if rowsAffected := tx.RowsAffected; rowsAffected == 0 {
		return errors.WithStack(persistence.ErrStaleObject)
	}

	return nil
}
