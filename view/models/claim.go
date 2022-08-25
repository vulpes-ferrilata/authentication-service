package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Claim struct {
	ID     primitive.ObjectID
	UserID primitive.ObjectID
}
