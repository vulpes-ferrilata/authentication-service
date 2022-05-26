package services

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/domain/mappers"
	"github.com/VulpesFerrilata/authentication-service/domain/models"
	"github.com/VulpesFerrilata/authentication-service/persistence/repositories"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserCredentialService interface {
	NewUserCredential(ctx context.Context, id uuid.UUID, userID uuid.UUID, email string, password string) (*models.UserCredential, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.UserCredential, error)
	GetByEmail(ctx context.Context, email string) (*models.UserCredential, error)
	Save(ctx context.Context, userCredential *models.UserCredential) error
	Delete(ctx context.Context, userCredential *models.UserCredential) error
}

func NewUserCredentialService(userCredentialRepository repositories.UserCredentialRepository,
	userCredentialMapper mappers.UserCredentialMapper) UserCredentialService {
	return &userCredentialService{
		userCredentialRepository: userCredentialRepository,
		userCredentialMapper:     userCredentialMapper,
	}
}

type userCredentialService struct {
	userCredentialRepository repositories.UserCredentialRepository
	userCredentialMapper     mappers.UserCredentialMapper
}

func (u userCredentialService) NewUserCredential(ctx context.Context, id uuid.UUID, userID uuid.UUID, email string, password string) (*models.UserCredential, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userCredential := models.NewUserCredential(id, userID, email, hashPassword)

	return userCredential, nil
}

func (u userCredentialService) GetByID(ctx context.Context, id uuid.UUID) (*models.UserCredential, error) {
	userCredentialEntity, err := u.userCredentialRepository.GetByID(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userCredential := u.userCredentialMapper.ToModel(ctx, userCredentialEntity)

	return userCredential, nil
}

func (u userCredentialService) GetByEmail(ctx context.Context, email string) (*models.UserCredential, error) {
	userCredentialEntity, err := u.userCredentialRepository.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userCredential := u.userCredentialMapper.ToModel(ctx, userCredentialEntity)

	return userCredential, nil
}

func (u userCredentialService) Save(ctx context.Context, userCredential *models.UserCredential) error {
	userCredentialEntity := u.userCredentialMapper.ToEntity(ctx, userCredential)

	if err := u.userCredentialRepository.Save(ctx, userCredentialEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (u userCredentialService) Delete(ctx context.Context, userCredential *models.UserCredential) error {
	userCredentialEntity := u.userCredentialMapper.ToEntity(ctx, userCredential)

	if err := u.userCredentialRepository.Delete(ctx, userCredentialEntity); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
