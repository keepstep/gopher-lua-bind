package modele

import (
	"errors"
)

type ModelE struct {
	Speed int
}

func (*ModelE) Run(speed int) (status int, err error) {
	return 1, errors.New("error")
}
