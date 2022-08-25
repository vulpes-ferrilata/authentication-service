package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Claim struct {
	aggregateRoot
	userID primitive.ObjectID
}

func (c Claim) GetUserID() primitive.ObjectID {
	return c.userID
}
