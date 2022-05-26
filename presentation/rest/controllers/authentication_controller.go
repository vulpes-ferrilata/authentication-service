package controllers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/commands"
	"github.com/VulpesFerrilata/authentication-service/application/queries"
	"github.com/VulpesFerrilata/authentication-service/application/queries/dtos"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/bus"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/saga"
	"github.com/VulpesFerrilata/authentication-service/presentation/rest/requests"
	"github.com/VulpesFerrilata/authentication-service/presentation/rest/responses"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/pkg/errors"
)

type AuthenticationController interface {
	PostLogin(ctx iris.Context) (mvc.Result, error)
	PostRefresh(ctx iris.Context) (mvc.Result, error)
	PostRevoke(ctx iris.Context) (mvc.Result, error)
}

func NewAuthenticationController(queryBus bus.QueryBus,
	commandBus bus.CommandBus) AuthenticationController {
	return &authenticationController{
		queryBus:   queryBus,
		commandBus: commandBus,
	}
}

type authenticationController struct {
	queryBus   bus.QueryBus
	commandBus bus.CommandBus
}

func (a authenticationController) PostLogin(ctx iris.Context) (mvc.Result, error) {
	loginRequest := new(requests.LoginRequest)

	if err := ctx.ReadJSON(loginRequest); err != nil {
		return nil, errors.WithStack(err)
	}

	getUserCredentialByEmailAndPasswordQuery := &queries.GetUserCredentialByEmailAndPasswordQuery{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
	result, err := a.queryBus.Execute(ctx.Request().Context(), getUserCredentialByEmailAndPasswordQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	userCredentialDTO := result.(*dtos.UserCredentialDTO)

	jti, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	createClaimCommand := &commands.CreateClaimCommand{
		UserID: userCredentialDTO.UserID,
		JTI:    jti.String(),
	}
	if err := a.commandBus.Execute(ctx.Request().Context(), createClaimCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	getAccessTokenByClaimQuery := &queries.GetAccessTokenByClaimQuery{
		UserID: userCredentialDTO.UserID,
		JTI:    jti.String(),
	}
	result, err = a.queryBus.Execute(ctx.Request().Context(), getAccessTokenByClaimQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	accessToken := result.(string)

	getRefreshTokenByClaimQuery := &queries.GetRefreshTokenByClaimQuery{
		UserID: userCredentialDTO.UserID,
		JTI:    jti.String(),
	}
	result, err = a.queryBus.Execute(ctx.Request().Context(), getRefreshTokenByClaimQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	refreshToken := result.(string)

	tokenResponse := &responses.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return &mvc.Response{
		Code:   iris.StatusOK,
		Object: tokenResponse,
	}, nil
}

func (a authenticationController) PostRegister(ctx iris.Context) (mvc.Result, error) {
	registerRequest := new(requests.RegisterRequest)

	if err := ctx.ReadJSON(registerRequest); err != nil {
		return nil, errors.WithStack(err)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	coordinator := saga.NewCoordinator()

	if err := coordinator.Execute(ctx.Request().Context(),
		&saga.Step{
			Handle: func(ctx context.Context) error {
				createUserCredentialCommand := &commands.CreateUserCredentialCommand{
					ID:       id.String(),
					UserID:   userID.String(),
					Email:    registerRequest.Email,
					Password: registerRequest.Password,
				}

				return a.commandBus.Execute(ctx, createUserCredentialCommand)
			},
			Compensate: func(ctx context.Context) error {
				removeUserCredentialCommand := &commands.RemoveUserCredentialCommand{
					ID: id.String(),
				}

				return a.commandBus.Execute(ctx, removeUserCredentialCommand)
			},
		},
		&saga.Step{
			Handle: func(ctx context.Context) error {
				createUserCommand := &commands.CreateUserCommand{
					ID:          userID.String(),
					DisplayName: registerRequest.DisplayName,
				}

				return a.commandBus.Execute(ctx, createUserCommand)
			},
		}); err != nil {
		return nil, errors.WithStack(err)
	}

	userResponse := &responses.UserResponse{
		ID: userID.String(),
	}

	return &mvc.Response{
		Code:   iris.StatusCreated,
		Object: userResponse,
	}, nil
}

func (a authenticationController) PostRefresh(ctx iris.Context) (mvc.Result, error) {
	refreshRequest := new(requests.RefreshRequest)

	if err := ctx.ReadJSON(refreshRequest); err != nil {
		return nil, errors.WithStack(err)
	}

	getClaimByRefreshTokenQuery := &queries.GetClaimByRefreshTokenQuery{
		RefreshToken: refreshRequest.RefreshToken,
	}
	result, err := a.queryBus.Execute(ctx.Request().Context(), getClaimByRefreshTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	claimDTO := result.(*dtos.ClaimDTO)

	getAccessTokenByClaimQuery := &queries.GetAccessTokenByClaimQuery{
		UserID: claimDTO.UserID,
		JTI:    claimDTO.JTI,
	}
	result, err = a.queryBus.Execute(ctx.Request().Context(), getAccessTokenByClaimQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	accessToken := result.(string)

	tokenResponse := &responses.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshRequest.RefreshToken,
	}

	return &mvc.Response{
		Code:   iris.StatusOK,
		Object: tokenResponse,
	}, nil
}

func (a authenticationController) PostRevoke(ctx iris.Context) (mvc.Result, error) {
	revokeRequest := new(requests.RevokeRequest)

	if err := ctx.ReadJSON(revokeRequest); err != nil {
		return nil, errors.WithStack(err)
	}

	getClaimByRefreshTokenQuery := &queries.GetClaimByRefreshTokenQuery{
		RefreshToken: revokeRequest.RefreshToken,
	}
	result, err := a.queryBus.Execute(ctx.Request().Context(), getClaimByRefreshTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	claimDTO := result.(*dtos.ClaimDTO)

	removeClaimCommand := &commands.RemoveClaimCommand{
		UserID: claimDTO.UserID,
		JTI:    claimDTO.JTI,
	}
	if err := a.commandBus.Execute(ctx.Request().Context(), removeClaimCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &mvc.Response{
		Code: iris.StatusOK,
	}, nil
}
