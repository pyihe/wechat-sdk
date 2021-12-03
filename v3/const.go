package v3

import "github.com/pyihe/go-pkg/errors"

var (
	ErrNoSecret    = errors.New("no app secret")
	ErrInvalidCode = errors.New("invalid code")
	ErrNotExistKey = errors.New("not exist key")
	ErrConvert     = errors.New("unsupported convert type")
)
