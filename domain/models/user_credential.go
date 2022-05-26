package models

import (
	"github.com/VulpesFerrilata/authentication-service/domain/models/common"
	"github.com/google/uuid"
)

func NewUserCredential(id uuid.UUID, userID uuid.UUID, email string, hashPassword []byte) *UserCredential {
	return &UserCredential{
		Entity:       common.NewEntity(id),
		userID:       userID,
		email:        email,
		hashPassword: hashPassword,
	}
}

type UserCredential struct {
	common.Entity
	userID       uuid.UUID
	email        string
	hashPassword []byte
}

func (u UserCredential) GetUserID() uuid.UUID {
	return u.userID
}

func (u UserCredential) GetEmail() string {
	return u.email
}

func (u UserCredential) GetHashPassword() []byte {
	return u.hashPassword
}
