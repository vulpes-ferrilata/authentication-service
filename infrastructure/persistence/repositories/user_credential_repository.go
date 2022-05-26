package repositories

import (
	"context"

	infrastructure_context "github.com/VulpesFerrilata/authentication-service/infrastructure/context"
	"github.com/VulpesFerrilata/authentication-service/persistence"
	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewUserCredentialRepository() repositories.UserCredentialRepository {
	return &userCredentialRepository{}
}

type userCredentialRepository struct {
	repository[entities.UserCredential]
}

func (u userCredentialRepository) IsEmailExists(ctx context.Context, email string) (bool, error) {
	count := int64(0)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return false, errors.WithStack(err)
	}

	tx = tx.Model(&entities.UserCredential{})
	tx = tx.Where("email = ?", email)
	tx = tx.Count(&count)
	if err := tx.Error; err != nil {
		return false, errors.WithStack(err)
	}

	return count > 0, nil
}

func (u userCredentialRepository) GetByEmail(ctx context.Context, email string) (*entities.UserCredential, error) {
	userCredentialEntity := new(entities.UserCredential)

	tx, err := infrastructure_context.GetTransaction(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tx = tx.Where("email = ?", email).First(userCredentialEntity)
	if err := tx.Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.WithStack(persistence.ErrRecordNotFound)
	}
	if err := tx.Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return userCredentialEntity, nil
}
