package entities

import "github.com/google/uuid"

type Claim struct {
	UserID uuid.UUID
	JTI    uuid.UUID
}
