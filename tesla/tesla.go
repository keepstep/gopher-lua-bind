package tesla

import (
	"errors"
	"fmt"
	"strings"

	"github.com/keepstep/gopher-lua-bind/tesla/aaaa"
	"github.com/keepstep/gopher-lua-bind/tesla/modele"
	"github.com/keepstep/gopher-lua-bind/tesla/modelx"
	"github.com/keepstep/gopher-lua-bind/tesla/modely"
)

type Cmp interface {
	Equal(name string) bool
}

type Tesla struct {
	Name    string
	ModelX  *modelx.ModelX
	ModelE  *modele.ModelE
	ModelY  *modely.ModelY
	Err     error
	Cars    []string
	Models  map[string]int
	RunCb   func(name string) string
	GoCb    func(name string, age float32) (string, int, bool)
	Compare Cmp
}

func (t *Tesla) Run(mode string, speed int) (status int, err error) {
	fmt.Printf("tesla.Run %s %d\n", mode, speed)
	if t.RunCb != nil {
		r := t.RunCb(mode)
		fmt.Printf("tesla.RunCb %s\n", r)
		// r = t.RunCb(mode + "02")
		// fmt.Printf("tesla.RunCb 02 %s\n", r)
		// r = t.RunCb(mode + "03")
		// fmt.Printf("tesla.RunCb 03 %s\n", r)
	}
	return 1, errors.New("error")
}
func (*Tesla) RunX(x *modelx.ModelX) (status int, err error) {
	return 1, errors.New("error")
}
func (t *Tesla) RunE() (status int, err error) {
	if t.ModelE != nil {
		t.ModelE.Run(0)
	}
	return 1, errors.New("error")
}

func (*Tesla) RunCmp(cmpa Cmp, cmpx modelx.CmpX) (status int, err error) {
	return 1, errors.New("error")
}

func (*Tesla) RunCallback(name string, ff func(name string, age int, flag []bool, mm map[string]int) (string, bool)) string {
	r, s := ff(strings.Repeat(name, 2), 100, []bool{true, false, true, false}, map[string]int{
		"name":   1,
		"age":    2,
		"gender": 3,
	})
	fmt.Printf("RunCallback %s,%t\n", r, s)
	return r
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

func TeslaCallBack(name string, ff func(name string, age int, flag bool) (string, bool)) string {
	r, _ := ff(strings.Repeat(name, 2), 100, true)
	return r
}

func TeslaCallInterface(name string, cmpa Cmp, cmpx modelx.CmpX) string {
	return ""
}

type CmpSSS struct {
	Name string
}

func (s *CmpSSS) Equal(name string) bool {
	fmt.Printf("CmpSSS Equal %s %s\n", s.Name, name)
	return s.Name == name
}

func TeslaGetCmp(name string) Cmp {
	fmt.Printf("TeslaGetCmp %s\n", name)
	return &CmpSSS{
		Name: name,
	}
}
