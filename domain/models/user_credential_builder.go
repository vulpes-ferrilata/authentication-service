package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCredentialBuilder struct {
	id           primitive.ObjectID
	userID       primitive.ObjectID
	email        string
	hashPassword []byte
	version      int
}

func (u UserCredentialBuilder) SetID(id primitive.ObjectID) UserCredentialBuilder {
	u.id = id

	return u
}

func (u UserCredentialBuilder) SetUserID(userID primitive.ObjectID) UserCredentialBuilder {
	u.userID = userID

	return u
}

func (u UserCredentialBuilder) SetEmail(email string) UserCredentialBuilder {
	u.email = email

	return u
}

func (u UserCredentialBuilder) SetHashPassword(hashPassword []byte) UserCredentialBuilder {
	u.hashPassword = hashPassword

	return u
}

func (u UserCredentialBuilder) SetVersion(version int) UserCredentialBuilder {
	u.version = version

	return u
}

func (u UserCredentialBuilder) Create() *UserCredential {
	return &UserCredential{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: u.id,
			},
			version: u.version,
		},
		userID:       u.userID,
		email:        u.email,
		hashPassword: u.hashPassword,
	}
}
