package models

import (
	"github.com/google/uuid"
)

func NewClaim(userID uuid.UUID, jti uuid.UUID) *Claim {
	return &Claim{
		userID: userID,
		jti:    jti,
	}
}

type Claim struct {
	userID uuid.UUID
	jti    uuid.UUID
}

func (c Claim) GetUserID() uuid.UUID {
	return c.userID
}

func (c Claim) GetJTI() uuid.UUID {
	return c.jti
}
