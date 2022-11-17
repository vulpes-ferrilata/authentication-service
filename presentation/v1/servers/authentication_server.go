package servers

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service-proto/pb"
	"github.com/vulpes-ferrilata/authentication-service-proto/pb/requests"
	"github.com/vulpes-ferrilata/authentication-service-proto/pb/responses"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/application/queries"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/command"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/cqrs/query"
	"github.com/vulpes-ferrilata/authentication-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAuthenticationServer(getTokenByClaimIDQueryHandler query.QueryHandler[*queries.GetTokenByClaimID, *models.Token],
	getClaimByAccessTokenQueryHandler query.QueryHandler[*queries.GetClaimByAccessToken, *models.Claim],
	getTokenByRefreshTokenQueryHandler query.QueryHandler[*queries.GetTokenByRefreshToken, *models.Token],
	createUserCredentialCommandHandler command.CommandHandler[*commands.CreateUserCredential],
	deleteUserCredentialCommandHandler command.CommandHandler[*commands.DeleteUserCredential],
	loginCommandHandler command.CommandHandler[*commands.Login],
	revokeTokenCommandHandler command.CommandHandler[*commands.RevokeToken]) pb.AuthenticationServer {
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
	pb.UnimplementedAuthenticationServer
	getTokenByClaimIDQueryHandler      query.QueryHandler[*queries.GetTokenByClaimID, *models.Token]
	getClaimByAccessTokenQueryHandler  query.QueryHandler[*queries.GetClaimByAccessToken, *models.Claim]
	getTokenByRefreshTokenQueryHandler query.QueryHandler[*queries.GetTokenByRefreshToken, *models.Token]
	createUserCredentialCommandHandler command.CommandHandler[*commands.CreateUserCredential]
	deleteUserCredentialCommandHandler command.CommandHandler[*commands.DeleteUserCredential]
	loginCommandHandler                command.CommandHandler[*commands.Login]
	revokeTokenCommandHandler          command.CommandHandler[*commands.RevokeToken]
}

func (a authenticationServer) GetTokenByClaimID(ctx context.Context, getTokenByClaimIDRequest *requests.GetTokenByClaimID) (*responses.Token, error) {
	getTokenByClaimIDQuery := &queries.GetTokenByClaimID{
		ClaimID: getTokenByClaimIDRequest.GetClaimID(),
	}

	token, err := a.getTokenByClaimIDQueryHandler.Handle(ctx, getTokenByClaimIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tokenResponse, err := mappers.TokenMapper{}.ToResponse(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tokenResponse, nil
}

func (a authenticationServer) GetClaimByAccessToken(ctx context.Context, getClaimByAccessTokenRequest *requests.GetClaimByAccessToken) (*responses.Claim, error) {
	getClaimByAccessTokenQuery := &queries.GetClaimByAccessToken{
		AccessToken: getClaimByAccessTokenRequest.GetAccessToken(),
	}

	claim, err := a.getClaimByAccessTokenQueryHandler.Handle(ctx, getClaimByAccessTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claimResponse, err := mappers.ClaimMapper{}.ToResponse(claim)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return claimResponse, nil
}

func (a authenticationServer) GetTokenByRefreshToken(ctx context.Context, getTokenByRefreshTokenRequest *requests.GetTokenByRefreshToken) (*responses.Token, error) {
	getTokenByRefreshTokenQuery := &queries.GetTokenByRefreshToken{
		RefreshToken: getTokenByRefreshTokenRequest.GetRefreshToken(),
	}

	token, err := a.getTokenByRefreshTokenQueryHandler.Handle(ctx, getTokenByRefreshTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tokenResponse, err := mappers.TokenMapper{}.ToResponse(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tokenResponse, nil
}

func (a authenticationServer) CreateUserCredential(ctx context.Context, createUserCredentialRequest *requests.CreateUserCredential) (*emptypb.Empty, error) {
	createUserCredentialCommand := &commands.CreateUserCredential{
		UserCredentialID: createUserCredentialRequest.GetUserCredentialID(),
		UserID:           createUserCredentialRequest.GetUserID(),
		Email:            createUserCredentialRequest.GetEmail(),
		Password:         createUserCredentialRequest.GetPassword(),
	}

	if err := a.createUserCredentialCommandHandler.Handle(ctx, createUserCredentialCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) DeleteUserCredential(ctx context.Context, deleteUserCredentialRequest *requests.DeleteUserCredential) (*emptypb.Empty, error) {
	deleteUserCredentialCommand := &commands.DeleteUserCredential{
		UserCredentialID: deleteUserCredentialRequest.GetUserCredentialID(),
	}

	if err := a.deleteUserCredentialCommandHandler.Handle(ctx, deleteUserCredentialCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) Login(ctx context.Context, loginRequest *requests.Login) (*emptypb.Empty, error) {
	loginCommand := &commands.Login{
		ClaimID:  loginRequest.GetClaimID(),
		Email:    loginRequest.GetEmail(),
		Password: loginRequest.GetPassword(),
	}

	if err := a.loginCommandHandler.Handle(ctx, loginCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) RevokeToken(ctx context.Context, revokeTokenRequest *requests.RevokeToken) (*emptypb.Empty, error) {
	revokeTokenCommand := &commands.RevokeToken{
		RefreshToken: revokeTokenRequest.GetRefreshToken(),
	}

	if err := a.revokeTokenCommandHandler.Handle(ctx, revokeTokenCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
