package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID primitive.ObjectID `bson:"_id"`
}
