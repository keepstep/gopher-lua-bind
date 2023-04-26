package modelx

import (
	"errors"
)

type ModelX struct {
	Speed int
}

func (*ModelX) Run(speed int) (status int, err error) {
	return 1, errors.New("error")
}
