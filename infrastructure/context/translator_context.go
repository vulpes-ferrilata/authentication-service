package context

import (
	"context"

	ut "github.com/go-playground/universal-translator"
	"github.com/pkg/errors"
)

type translatorContextKey struct{}

func WithTranslator(ctx context.Context, translator ut.Translator) context.Context {
	return context.WithValue(ctx, translatorContextKey{}, translator)
}

func GetTranslator(ctx context.Context) (ut.Translator, error) {
	translator, ok := ctx.Value(translatorContextKey{}).(ut.Translator)
	if !ok {
		return nil, errors.Wrap(ErrContextValueNotFound, "translator")
	}

	return translator, nil
}
