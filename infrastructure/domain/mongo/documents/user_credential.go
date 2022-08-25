package documents

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCredential struct {
	DocumentRoot `bson:",inline"`
	UserID       primitive.ObjectID `bson:"user_id"`
	Email        string             `bson:"email"`
	HashPassword []byte             `bson:"hash_password"`
	Version      int                `bson:"version"`
}
