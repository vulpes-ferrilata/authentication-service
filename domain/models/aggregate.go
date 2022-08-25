package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type aggregate struct {
	id primitive.ObjectID
}

func (a aggregate) GetID() primitive.ObjectID {
	return a.id
}
