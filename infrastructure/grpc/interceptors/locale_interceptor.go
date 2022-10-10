package interceptors

import (
	"context"
	"strings"

	httpext "github.com/go-playground/pkg/v5/net/http"
	"github.com/vulpes-ferrilata/authentication-service/infrastructure/context_values"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func NewLocaleInterceptor() *LocaleInterceptor {
	return &LocaleInterceptor{}
}

type LocaleInterceptor struct{}

func (l LocaleInterceptor) ServerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		locales := md[strings.ToLower(httpext.AcceptedLanguage)]
		ctx = context_values.WithLocales(ctx, locales)

		return handler(ctx, req)
	}
}
