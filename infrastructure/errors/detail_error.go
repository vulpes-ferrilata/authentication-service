package errors

import ut "github.com/go-playground/universal-translator"

type DetailError interface {
	error
	Translate(translator ut.Translator) (string, error)
}
