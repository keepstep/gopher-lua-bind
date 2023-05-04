package modelx

import (
	"errors"
)

type CmpX interface {
	Equal(name string) bool
}

type ModelX struct {
	Speed int
	CmpX  CmpX
}

func (*ModelX) Run(speed int) (status int, err error) {
	return 1, errors.New("error")
}
