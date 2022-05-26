package rest

import (
	"github.com/VulpesFerrilata/authentication-service/infrastructure/middlewares"
	"github.com/VulpesFerrilata/authentication-service/presentation/rest/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type Router interface {
	Init(app *iris.Application)
}

func NewRouter(translatorMiddleware *middlewares.TranslatorMiddleware,
	errorHandlerMiddleware *middlewares.ErrorHandlerMiddleware,
	authenticationController controllers.AuthenticationController) Router {
	return &router{
		translatorMiddleware:     translatorMiddleware,
		errorHandlerMiddleware:   errorHandlerMiddleware,
		authenticationController: authenticationController,
	}
}

type router struct {
	translatorMiddleware     *middlewares.TranslatorMiddleware
	errorHandlerMiddleware   *middlewares.ErrorHandlerMiddleware
	authenticationController controllers.AuthenticationController
}

func (r router) Init(app *iris.Application) {
	api := app.Party("/api")

	authApi := api.Party("/auth")
	authApi.Use(r.translatorMiddleware.Serve)
	authMvc := mvc.New(authApi)
	authMvc.HandleError(r.errorHandlerMiddleware.Handle)
	authMvc.Handle(r.authenticationController)
}
