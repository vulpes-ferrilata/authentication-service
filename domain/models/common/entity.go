package common

import "github.com/google/uuid"

func NewEntity(id uuid.UUID) Entity {
	entity := Entity{}
	entity.id = id
	return entity
}

type Entity struct {
	id uuid.UUID
}

func (e Entity) GetID() uuid.UUID {
	return e.id
}
