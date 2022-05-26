package models

import (
	"github.com/VulpesFerrilata/authentication-service/domain/models/common"
	"github.com/google/uuid"
)

func NewUser(id uuid.UUID, displayName string) *User {
	return &User{
		Entity:      common.NewEntity(id),
		displayName: displayName,
	}
}

type User struct {
	common.Entity
	displayName string
}

func (u User) GetDisplayName() string {
	return u.displayName
}
