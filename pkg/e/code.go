package e

import "errors"

var (
	ErrorNilParam   = errors.New("check param")
	ErrorNoOpenId   = errors.New("check param: openid")
	ErrorNoToken    = errors.New("check param: access_token")
	ErrorInitClinet = errors.New("call NewClientWithParam first")
)

const (
	MALE   = 1
	FEMALE = 2
)
