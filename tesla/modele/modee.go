package modele

import (
	"errors"
	"fmt"
)

type ModelE struct {
	Speed int
}

func (e *ModelE) Run(speed int) (status int, err error) {
	fmt.Printf("modele run %d %d\n", speed, e.Speed)
	return 1, errors.New("error")
}
