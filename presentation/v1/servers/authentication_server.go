package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/application/queries"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/authentication-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/shared/proto/v1/authentication"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAuthenticationServer(getTokenByClaimIDQueryHandler query.QueryHandler[*queries.GetTokenByClaimIDQuery, *models.Token],
	getClaimByAccessTokenQueryHandler query.QueryHandler[*queries.GetClaimByAccessTokenQuery, *models.Claim],
	getTokenByRefreshTokenQueryHandler query.QueryHandler[*queries.GetTokenByRefreshTokenQuery, *models.Token],
	createUserCredentialCommandHandler command.CommandHandler[*commands.CreateUserCredentialCommand],
	deleteUserCredentialCommandHandler command.CommandHandler[*commands.DeleteUserCredentialCommand],
	loginCommandHandler command.CommandHandler[*commands.LoginCommand],
	revokeTokenCommandHandler command.CommandHandler[*commands.RevokeTokenCommand]) authentication.AuthenticationServer {
	return &authenticationServer{
		getTokenByClaimIDQueryHandler:      getTokenByClaimIDQueryHandler,
		getClaimByAccessTokenQueryHandler:  getClaimByAccessTokenQueryHandler,
		getTokenByRefreshTokenQueryHandler: getTokenByRefreshTokenQueryHandler,
		createUserCredentialCommandHandler: createUserCredentialCommandHandler,
		deleteUserCredentialCommandHandler: deleteUserCredentialCommandHandler,
		loginCommandHandler:                loginCommandHandler,
		revokeTokenCommandHandler:          revokeTokenCommandHandler,
	}
}

type authenticationServer struct {
	authentication.UnimplementedAuthenticationServer
	getTokenByClaimIDQueryHandler      query.QueryHandler[*queries.GetTokenByClaimIDQuery, *models.Token]
	getClaimByAccessTokenQueryHandler  query.QueryHandler[*queries.GetClaimByAccessTokenQuery, *models.Claim]
	getTokenByRefreshTokenQueryHandler query.QueryHandler[*queries.GetTokenByRefreshTokenQuery, *models.Token]
	createUserCredentialCommandHandler command.CommandHandler[*commands.CreateUserCredentialCommand]
	deleteUserCredentialCommandHandler command.CommandHandler[*commands.DeleteUserCredentialCommand]
	loginCommandHandler                command.CommandHandler[*commands.LoginCommand]
	revokeTokenCommandHandler          command.CommandHandler[*commands.RevokeTokenCommand]
}

func (a authenticationServer) GetTokenByClaimID(ctx context.Context, getTokenByClaimIDRequest *authentication.GetTokenByClaimIDRequest) (*authentication.TokenResponse, error) {
	getTokenByClaimIDQuery := &queries.GetTokenByClaimIDQuery{
		ClaimID: getTokenByClaimIDRequest.GetClaimID(),
	}

	token, err := a.getTokenByClaimIDQueryHandler.Handle(ctx, getTokenByClaimIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tokenResponse := mappers.ToTokenResponse(token)

	return tokenResponse, nil
}

func (a authenticationServer) GetClaimByAccessToken(ctx context.Context, getClaimByAccessTokenRequest *authentication.GetClaimByAccessTokenRequest) (*authentication.ClaimResponse, error) {
	getClaimByAccessTokenQuery := &queries.GetClaimByAccessTokenQuery{
		AccessToken: getClaimByAccessTokenRequest.GetAccessToken(),
	}

	claim, err := a.getClaimByAccessTokenQueryHandler.Handle(ctx, getClaimByAccessTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claimResponse := mappers.ToClaimResponse(claim)

	return claimResponse, nil
}

func (a authenticationServer) GetTokenByRefreshToken(ctx context.Context, getTokenByRefreshTokenRequest *authentication.GetTokenByRefreshTokenRequest) (*authentication.TokenResponse, error) {
	getTokenByRefreshTokenQuery := &queries.GetTokenByRefreshTokenQuery{
		RefreshToken: getTokenByRefreshTokenRequest.GetRefreshToken(),
	}

	token, err := a.getTokenByRefreshTokenQueryHandler.Handle(ctx, getTokenByRefreshTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tokenResponse := mappers.ToTokenResponse(token)

	return tokenResponse, nil
}

func (a authenticationServer) CreateUserCredential(ctx context.Context, createUserCredentialRequest *authentication.CreateUserCredentialRequest) (*emptypb.Empty, error) {
	createUserCredentialCommand := &commands.CreateUserCredentialCommand{
		ID:       createUserCredentialRequest.GetID(),
		UserID:   createUserCredentialRequest.GetUserID(),
		Email:    createUserCredentialRequest.GetEmail(),
		Password: createUserCredentialRequest.GetPassword(),
	}

	if err := a.createUserCredentialCommandHandler.Handle(ctx, createUserCredentialCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) DeleteUserCredential(ctx context.Context, deleteUserCredentialRequest *authentication.DeleteUserCredentialRequest) (*emptypb.Empty, error) {
	deleteUserCredentialCommand := &commands.DeleteUserCredentialCommand{
		ID: deleteUserCredentialRequest.GetID(),
	}

	if err := a.deleteUserCredentialCommandHandler.Handle(ctx, deleteUserCredentialCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) Login(ctx context.Context, loginRequest *authentication.LoginRequest) (*emptypb.Empty, error) {
	loginCommand := &commands.LoginCommand{
		ClaimID:  loginRequest.GetClaimID(),
		Email:    loginRequest.GetEmail(),
		Password: loginRequest.GetPassword(),
	}

	if err := a.loginCommandHandler.Handle(ctx, loginCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) RevokeToken(ctx context.Context, revokeTokenRequest *authentication.RevokeTokenRequest) (*emptypb.Empty, error) {
	revokeTokenCommand := &commands.RevokeTokenCommand{
		RefreshToken: revokeTokenRequest.GetRefreshToken(),
	}

	if err := a.revokeTokenCommandHandler.Handle(ctx, revokeTokenCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
