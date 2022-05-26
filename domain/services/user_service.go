package services

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/domain/models"
	proto_user "github.com/VulpesFerrilata/shared/proto/user"
	"github.com/asim/go-micro/v3/client"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type UserService interface {
	NewUser(ctx context.Context, id uuid.UUID, displayName string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}

func NewUserService(client client.Client) UserService {
	return &userService{
		userService: proto_user.NewUserService("boardgame.user.service", client),
	}
}

type userService struct {
	userService proto_user.UserService
}

func (u userService) NewUser(ctx context.Context, id uuid.UUID, displayName string) (*models.User, error) {
	user := models.NewUser(id, displayName)
	return user, nil
}

func (u userService) Create(ctx context.Context, user *models.User) error {
	createUserRequest := new(proto_user.CreateUserRequest)
	createUserRequest.ID = user.GetID().String()
	createUserRequest.DisplayName = user.GetDisplayName()

	if _, err := u.userService.CreateUser(ctx, createUserRequest); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
