package models

import (
	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/app_errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserCredential struct {
	aggregateRoot
	userID       primitive.ObjectID
	email        string
	hashPassword []byte
}

func (u UserCredential) GetUserID() primitive.ObjectID {
	return u.userID
}

func (u UserCredential) GetEmail() string {
	return u.email
}

func (u UserCredential) GetHashPassword() []byte {
	return u.hashPassword
}

func (u *UserCredential) SetPassword(password string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(app_errors.ErrUnableToEncryptPassword)
	}
	u.hashPassword = hashPassword

	return nil
}

func (u UserCredential) ComparePassword(password string) error {
	if err := bcrypt.CompareHashAndPassword(u.hashPassword, []byte(password)); err != nil {
		return errors.WithStack(app_errors.ErrPasswordIsInvalid)
	}

	return nil
}
