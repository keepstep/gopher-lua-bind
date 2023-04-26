package tesla

import (
	"errors"
	"fmt"

	"github.com/keepstep/gopher-lua-bind/tesla/aaaa"
	"github.com/keepstep/gopher-lua-bind/tesla/modele"
	"github.com/keepstep/gopher-lua-bind/tesla/modelx"
	"github.com/keepstep/gopher-lua-bind/tesla/modely"
)

type Tesla struct {
	Name   string
	ModelX *modelx.ModelX
	ModelE *modele.ModelE
	ModelY *modely.ModelY
	Err    error
	Cars   []string
	Models map[string]int
}

func (*Tesla) Run(mode string, speed int) (status int, err error) {
	return 1, errors.New("error")
}
func (*Tesla) RunX(x *modelx.ModelX) (status int, err error) {
	return 1, errors.New("error")
}
func (*Tesla) RunE(x *modele.ModelE) (status int, err error) {
	return 1, errors.New("error")
}

func TeslaCompare(str string) int {
	return 500
}

func TeslaTest(a int, b float32, m map[string]int) (r int, f float64, ss []string, err error) {
	ss = []string{}
	for _, v := range m {
		ss = append(ss, fmt.Sprintf("-x-%d-x-", v))
	}
	return a * 2, float64(b * 2), ss, errors.New("error tesla")
}

func TeslaGetAAA(name string) *aaaa.AAA {
	return &aaaa.AAA{Name: "TeslaAAA", Age: 200}
}
