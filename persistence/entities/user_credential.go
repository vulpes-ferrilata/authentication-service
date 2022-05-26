package entities

import (
	"github.com/VulpesFerrilata/authentication-service/persistence/entities/common"
	"github.com/google/uuid"
)

type UserCredential struct {
	common.Entity
	UserID       uuid.UUID `gorm:"unique"`
	Email        string    `gorm:"type:varchar(20); unique"`
	HashPassword []byte
}
