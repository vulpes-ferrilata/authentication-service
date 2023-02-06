package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vulpes-ferrilata/authentication-service-proto/pb"
	pb_models "github.com/vulpes-ferrilata/authentication-service-proto/pb/models"
	"github.com/vulpes-ferrilata/authentication-service/application/commands"
	"github.com/vulpes-ferrilata/authentication-service/application/queries"
	"github.com/vulpes-ferrilata/authentication-service/presentation/v1/mappers"
	"github.com/vulpes-ferrilata/authentication-service/view/models"
	"github.com/vulpes-ferrilata/cqrs"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewAuthenticationServer(queryBus *cqrs.QueryBus,
	commandBus *cqrs.CommandBus) pb.AuthenticationServer {
	return &authenticationServer{
		queryBus:   queryBus,
		commandBus: commandBus,
	}
}

type authenticationServer struct {
	pb.UnimplementedAuthenticationServer
	queryBus   *cqrs.QueryBus
	commandBus *cqrs.CommandBus
}

func (a authenticationServer) GetTokenByClaimID(ctx context.Context, getTokenByClaimIDRequest *pb_models.GetTokenByClaimIDRequest) (*pb_models.Token, error) {
	getTokenByClaimIDQuery := &queries.GetTokenByClaimIDQuery{
		ClaimID: getTokenByClaimIDRequest.GetClaimID(),
	}

	token, err := cqrs.ParseQueryHandlerFunc[*queries.GetTokenByClaimIDQuery, *models.Token](a.queryBus.Execute)(ctx, getTokenByClaimIDQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tokenResponse, err := mappers.TokenMapper{}.ToResponse(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tokenResponse, nil
}

func (a authenticationServer) GetClaimByAccessToken(ctx context.Context, getClaimByAccessTokenRequest *pb_models.GetClaimByAccessTokenRequest) (*pb_models.Claim, error) {
	getClaimByAccessTokenQuery := &queries.GetClaimByAccessTokenQuery{
		AccessToken: getClaimByAccessTokenRequest.GetAccessToken(),
	}

	claim, err := cqrs.ParseQueryHandlerFunc[*queries.GetClaimByAccessTokenQuery, *models.Claim](a.queryBus.Execute)(ctx, getClaimByAccessTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	claimResponse, err := mappers.ClaimMapper{}.ToResponse(claim)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return claimResponse, nil
}

func (a authenticationServer) GetTokenByRefreshToken(ctx context.Context, getTokenByRefreshTokenRequest *pb_models.GetTokenByRefreshTokenRequest) (*pb_models.Token, error) {
	getTokenByRefreshTokenQuery := &queries.GetTokenByRefreshTokenQuery{
		RefreshToken: getTokenByRefreshTokenRequest.GetRefreshToken(),
	}

	token, err := cqrs.ParseQueryHandlerFunc[*queries.GetTokenByRefreshTokenQuery, *models.Token](a.queryBus.Execute)(ctx, getTokenByRefreshTokenQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tokenResponse, err := mappers.TokenMapper{}.ToResponse(token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tokenResponse, nil
}

func (a authenticationServer) CreateUserCredential(ctx context.Context, createUserCredentialRequest *pb_models.CreateUserCredentialRequest) (*emptypb.Empty, error) {
	createUserCredentialCommand := &commands.CreateUserCredentialCommand{
		UserCredentialID: createUserCredentialRequest.GetUserCredentialID(),
		UserID:           createUserCredentialRequest.GetUserID(),
		Email:            createUserCredentialRequest.GetEmail(),
		Password:         createUserCredentialRequest.GetPassword(),
	}

	if err := a.commandBus.Execute(ctx, createUserCredentialCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) DeleteUserCredential(ctx context.Context, deleteUserCredentialRequest *pb_models.DeleteUserCredentialRequest) (*emptypb.Empty, error) {
	deleteUserCredentialCommand := &commands.DeleteUserCredentialCommand{
		UserCredentialID: deleteUserCredentialRequest.GetUserCredentialID(),
	}

	if err := a.commandBus.Execute(ctx, deleteUserCredentialCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) Login(ctx context.Context, loginRequest *pb_models.LoginRequest) (*emptypb.Empty, error) {
	loginCommand := &commands.LoginCommand{
		ClaimID:  loginRequest.GetClaimID(),
		Email:    loginRequest.GetEmail(),
		Password: loginRequest.GetPassword(),
	}

	if err := a.commandBus.Execute(ctx, loginCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}

func (a authenticationServer) RevokeToken(ctx context.Context, revokeTokenRequest *pb_models.RevokeTokenRequest) (*emptypb.Empty, error) {
	revokeTokenCommand := &commands.RevokeTokenCommand{
		RefreshToken: revokeTokenRequest.GetRefreshToken(),
	}

	if err := a.commandBus.Execute(ctx, revokeTokenCommand); err != nil {
		return nil, errors.WithStack(err)
	}

	return &emptypb.Empty{}, nil
}
