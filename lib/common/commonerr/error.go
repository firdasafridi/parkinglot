package commonerr

import "errors"

var (
	ErrCommon = errors.New("failure occured")
)

func commonErr() bool {
	return true
}
