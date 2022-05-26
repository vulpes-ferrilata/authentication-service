package middlewares

import (
	"context"

	infrastructure_context "github.com/VulpesFerrilata/authentication-service/infrastructure/context"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	httpext "github.com/go-playground/pkg/v5/net/http"
	"github.com/go-playground/pure"
	ut "github.com/go-playground/universal-translator"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
)

func NewTranslatorMiddleware(universalTranslator *ut.UniversalTranslator) *TranslatorMiddleware {
	return &TranslatorMiddleware{
		universalTranslator: universalTranslator,
	}
}

type TranslatorMiddleware struct {
	universalTranslator *ut.UniversalTranslator
}

func (t TranslatorMiddleware) WrapCall(f client.CallFunc) client.CallFunc {
	return func(ctx context.Context, node *registry.Node, request client.Request, response interface{}, opts client.CallOptions) error {
		translator, err := infrastructure_context.GetTranslator(ctx)
		if err != nil {
			return errors.WithStack(err)
		}
		ctx = metadata.Set(ctx, httpext.AcceptedLanguage, translator.Locale())
		return f(ctx, node, request, response, opts)
	}
}

func (t TranslatorMiddleware) WrapHandler(f server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, request server.Request, response interface{}) error {
		language, _ := metadata.Get(ctx, httpext.AcceptedLanguage)
		translator, _ := t.universalTranslator.FindTranslator(language)
		ctx = infrastructure_context.WithTranslator(ctx, translator)
		return f(ctx, request, response)
	}
}

func (t TranslatorMiddleware) Serve(ctx iris.Context) {
	request := ctx.Request()
	requestCtx := request.Context()
	languages := pure.AcceptedLanguages(request)
	translator, _ := t.universalTranslator.FindTranslator(languages...)

	requestCtx = infrastructure_context.WithTranslator(requestCtx, translator)
	request = request.WithContext(requestCtx)
	ctx.ResetRequest(request)

	ctx.Next()
}
