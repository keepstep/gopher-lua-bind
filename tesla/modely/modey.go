package modely

import (
	"errors"
)

type ModelY struct {
	Speed int
}

func (*ModelY) Run(speed int) (status int, err error) {
	return 1, errors.New("error")
}
