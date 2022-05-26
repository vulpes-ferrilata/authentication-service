package handlers

import (
	"context"

	"github.com/VulpesFerrilata/authentication-service/application/queries"
	"github.com/VulpesFerrilata/authentication-service/application/queries/dtos"
	domain_services "github.com/VulpesFerrilata/authentication-service/domain/services"
	"github.com/VulpesFerrilata/authentication-service/domain/validators"
	"github.com/VulpesFerrilata/authentication-service/infrastructure/dig/results"
	"github.com/pkg/errors"
)

func NewGetUserCredentialByEmailAndPasswordQueryHandler(userCredentialValidator validators.UserCredentialValidator,
	userCredentialService domain_services.UserCredentialService) results.QueryHandlerResult {
	queryHandler := &getUserCredentialByEmailAndPasswordQueryHandler{
		userCredentialValidator: userCredentialValidator,
		userCredentialService:   userCredentialService,
	}

	return results.QueryHandlerResult{
		QueryHandler: queryHandler,
	}
}

type getUserCredentialByEmailAndPasswordQueryHandler struct {
	userCredentialValidator validators.UserCredentialValidator
	userCredentialService   domain_services.UserCredentialService
}

func (g getUserCredentialByEmailAndPasswordQueryHandler) GetQuery() interface{} {
	return &queries.GetUserCredentialByEmailAndPasswordQuery{}
}

func (g getUserCredentialByEmailAndPasswordQueryHandler) Handle(ctx context.Context, query interface{}) (interface{}, error) {
	getUserCredentialByEmailAndPasswordQuery := query.(*queries.GetUserCredentialByEmailAndPasswordQuery)

	if err := g.userCredentialValidator.ValidateAuthentication(ctx, getUserCredentialByEmailAndPasswordQuery.Email, getUserCredentialByEmailAndPasswordQuery.Password); err != nil {
		return nil, errors.WithStack(err)
	}

	userCredential, err := g.userCredentialService.GetByEmail(ctx, getUserCredentialByEmailAndPasswordQuery.Email)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	userCredentialDTO := dtos.NewUserCredentialDTO(userCredential)

	return userCredentialDTO, nil
}
