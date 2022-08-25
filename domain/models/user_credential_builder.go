package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCredentialBuilder interface {
	SetID(id primitive.ObjectID) UserCredentialBuilder
	SetUserID(userID primitive.ObjectID) UserCredentialBuilder
	SetEmail(email string) UserCredentialBuilder
	SetHashPassword(hashPassword []byte) UserCredentialBuilder
	SetVersion(version int) UserCredentialBuilder
	Create() *UserCredential
}

func NewUserCredentialBuilder() UserCredentialBuilder {
	return &userCredentialBuilder{}
}

type userCredentialBuilder struct {
	id           primitive.ObjectID
	userID       primitive.ObjectID
	email        string
	hashPassword []byte
	version      int
}

func (u *userCredentialBuilder) SetID(id primitive.ObjectID) UserCredentialBuilder {
	u.id = id

	return u
}

func (u *userCredentialBuilder) SetUserID(userID primitive.ObjectID) UserCredentialBuilder {
	u.userID = userID

	return u
}

func (u *userCredentialBuilder) SetEmail(email string) UserCredentialBuilder {
	u.email = email

	return u
}

func (u *userCredentialBuilder) SetHashPassword(hashPassword []byte) UserCredentialBuilder {
	u.hashPassword = hashPassword

	return u
}

func (u *userCredentialBuilder) SetVersion(version int) UserCredentialBuilder {
	u.version = version

	return u
}

func (u userCredentialBuilder) Create() *UserCredential {
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
