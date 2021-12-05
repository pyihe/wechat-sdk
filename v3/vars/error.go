package vars

import "github.com/pyihe/go-pkg/errors"

var (
	ErrNoParam           = errors.New("no param.")
	ErrNoSecret          = errors.New("no app secret.")
	ErrNoSerialNo        = errors.New("no serial_no.")
	ErrNoMchId           = errors.New("no mchid.")
	ErrNoApiV3Key        = errors.New("no api v3 key.")
	ErrInvalidSessionKey = errors.New("cannot get \"session_key\".")
	ErrNotExistKey       = errors.New("not exist key.")
	ErrConvert           = errors.New("unsupported convert type.")
	ErrRequestAgain      = errors.New("please request once again.")
)
