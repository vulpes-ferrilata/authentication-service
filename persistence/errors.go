package persistence

import "github.com/pkg/errors"

var (
	ErrRecordNotFound error = errors.New("record not found")
	ErrStaleObject    error = errors.New("attemped to update or delete a stale object")
)
