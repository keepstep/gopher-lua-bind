package tesla

import (
	"fmt"

	"github.com/keepstep/gopher-lua-bind/car/base"
)

type Modely struct {
	Name      string
	Seats     uint
	Mpg       float32
	Price     int
	Sale      bool
	CurrSpeed int
	Lights    []*base.CarLight
	Doors     map[string]*base.CarDoor

	Metric map[string]int
	Intro  map[string]string
	Param  map[string]any
}

func (t *Modely) SpeedUp(add int) int {
	t.CurrSpeed += 2 * add
	return t.CurrSpeed
}
func (t *Modely) Park() bool {
	return true
}

func (t *Modely) Info() string {
	return fmt.Sprintf("Modely-%s-%d-%f", t.Name, t.Price, t.Mpg)
}
