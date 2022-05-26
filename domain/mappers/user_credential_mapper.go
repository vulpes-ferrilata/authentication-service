package mappers

import (
	"context"
	"sync"

	"github.com/VulpesFerrilata/authentication-service/domain/models"
	"github.com/VulpesFerrilata/authentication-service/persistence/entities"
)

type UserCredentialMapper interface {
	ToEntity(ctx context.Context, userCredential *models.UserCredential) *entities.UserCredential
	ToModel(ctx context.Context, userCredentialEntity *entities.UserCredential) *models.UserCredential
}

func NewUserCredentialMapper() UserCredentialMapper {
	return &userCredentialMapper{
		m: make(map[*models.UserCredential]*entities.UserCredential),
	}
}

type userCredentialMapper struct {
	m  map[*models.UserCredential]*entities.UserCredential
	mu sync.RWMutex
}

func (u *userCredentialMapper) ToEntity(ctx context.Context, userCredential *models.UserCredential) *entities.UserCredential {
	if userCredential == nil {
		return nil
	}

	u.mu.RLock()
	userCredentialEntity, ok := u.m[userCredential]
	u.mu.RUnlock()
	if !ok {
		userCredentialEntity = new(entities.UserCredential)

		u.mu.Lock()
		u.m[userCredential] = userCredentialEntity
		u.mu.Unlock()

		go func(userCredential *models.UserCredential, done <-chan struct{}) {
			<-done
			u.mu.Lock()
			delete(u.m, userCredential)
			u.mu.Unlock()
		}(userCredential, ctx.Done())
	}

	userCredentialEntity.ID = userCredential.GetID()
	userCredentialEntity.UserID = userCredential.GetUserID()
	userCredentialEntity.Email = userCredential.GetEmail()
	userCredentialEntity.HashPassword = userCredential.GetHashPassword()

	return userCredentialEntity
}

func (u *userCredentialMapper) ToModel(ctx context.Context, userCredentialEntity *entities.UserCredential) *models.UserCredential {
	if userCredentialEntity == nil {
		return nil
	}

	userCredential := models.NewUserCredential(
		userCredentialEntity.ID,
		userCredentialEntity.UserID,
		userCredentialEntity.Email,
		userCredentialEntity.HashPassword,
	)

	u.mu.Lock()
	u.m[userCredential] = userCredentialEntity
	u.mu.Unlock()

	go func(userCredential *models.UserCredential, done <-chan struct{}) {
		<-done
		u.mu.Lock()
		delete(u.m, userCredential)
		u.mu.Unlock()
	}(userCredential, ctx.Done())

	return userCredential
}
