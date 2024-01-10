package car

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/keepstep/gopher-lua-bind/car/base"
	"github.com/keepstep/gopher-lua-bind/car/byd"
	"github.com/keepstep/gopher-lua-bind/car/tesla"
)

type Fight interface {
	Fight(name string) (string, error)
}

type Driver struct {
	Name     string
	Age      uint16
	Han      *byd.Han
	Mdy      *tesla.Modely
	Byds     []*byd.Han
	Teslas   map[string]*tesla.Modely
	UsingCar base.CarBase
	FunStart func(car base.CarBase, speed int) (bool, error, *byd.Han, base.CarBase)
	FunStop  func([]Fight, map[string]Fight, []*byd.Han, map[int]*tesla.Modely, []string, map[string]float32) (bool, int, float32, string, []string, []int, map[int]float64, []Fight, map[string]Fight, map[int]*tesla.Modely, []*byd.Han)
}

func (t *Driver) Init() {
	t.Byds = []*byd.Han{}
	t.Teslas = map[string]*tesla.Modely{}
}

func (t *Driver) Fight(f Fight) {
	f.Fight("fight")
}

func (t *Driver) Drive(car base.CarBase, check func(second int) bool) error {
	t.UsingCar = car
	s := 0
	for {
		if check(s) {
			s++
			car.SpeedUp(1)
		} else {
			break
		}
	}
	car.Park()
	return errors.New("perfect")
}

func (t *Driver) Compare(a *byd.Han, b *tesla.Modely) bool {
	return a.CurrSpeed > b.CurrSpeed
}

func (t *Driver) BuyByd(a []*byd.Han) int {
	for _, c := range a {
		fmt.Println(c.Info())
		t.Byds = append(t.Byds, c)
	}
	return len(a)
}

func (t *Driver) BuyTesla(a map[string]*tesla.Modely) []string {
	r := []string{}
	for _, c := range a {
		fmt.Println(c.Info())
		t.Teslas[c.Name] = c
		r = append(r, c.Name)
	}
	return r
}

func (t *Driver) DoA(a int, p *byd.Han, c *tesla.Modely) (name string, aaa *byd.Han, bbb *tesla.Modely, err error) {
	return "name", &byd.Han{}, nil, errors.New("error do")
}

func (t *Driver) DoB(a int, p *byd.Han, h *http.Client) (name string, err error) {
	return "", nil
}

func (t *Driver) DoAdd(a int, b float32) (r int, err error) {
	return a + int(b), errors.New("add_no_error")
}

func (t *Driver) DoMap(a int, b float32, m map[string]int) (r int, f float64, mm map[string]string, err error) {
	mm = map[string]string{}
	for k, v := range m {
		mm[k] = fmt.Sprintf("-%d-", v)
	}
	return a * 2, float64(b * 2), mm, errors.New("map_no_error")
}

func (t *Driver) DoSlice(s []int) (ss []string, err error) {
	ss = []string{}
	for _, v := range s {
		ss = append(ss, fmt.Sprintf("-%d-", v))
	}
	return ss, errors.New("slice_no_error")
}

func (t *Driver) DoSliceSlice(s []int) (ss [][]string, dd [][]float32, mm []map[string]int, err error) {
	ss = [][]string{}
	dd = [][]float32{}
	mm = []map[string]int{}
	return ss, dd, mm, errors.New("slice_slice_no_error")
}

func (t *Driver) DoMapMap(a int, b float32, m map[string]int) (r int, f float64, mm map[string]map[string]int, ms map[string][]string, err error) {
	mm = map[string]map[string]int{}
	ms = map[string][]string{}
	return a * 2, float64(b * 2), mm, ms, errors.New("map_map_no_error")
}

func (t *Driver) DoMapAny(a int, b any, m map[string]any) (any, map[string]any) {
	bs, _ := json.Marshal(m)
	return 100, map[string]any{
		"123":  123,
		"abc":  "cba",
		"any1": a,
		"any2": b,
		"any3": string(bs),
	}
}

func (t *Driver) DoCb(name string, mm map[string]int, ss []int, ff func(name string, age int, flag []bool, mm map[string]int) (string, bool)) string {
	r, s := ff(strings.Repeat(name, 2), 100, []bool{true, false, true, false}, map[string]int{
		"name":   1,
		"age":    2,
		"gender": 3,
	})
	fmt.Printf("DoCb %s,%t\n", r, s)
	return r
}

func (t *Driver) DoCb2(ff func(han *byd.Han, fight Fight) (string, bool)) string {
	return ""
}

func GetByCb(name string, ff func(name string, age int, flag bool) (string, bool)) string {
	r, _ := ff(strings.Repeat(name, 3), 100, true)
	return r
}

func GetBrand(madein string, engine int, speed float64) (s []float64, o map[string]string) {
	return []float64{1.1 + speed, 1.2 + speed, 1.3 + speed, 1.4 + speed}, map[string]string{"byd": "C" + madein, "toyota": "J" + madein, "honda": "J" + madein, "bba": "G" + madein}
}

func GetCars(madein string, engine int, speed float32) (s []*byd.Han, o map[int]*tesla.Modely) {
	return []*byd.Han{
			{
				Name:  "H001",
				Price: 99,
			},
			{
				Name:  "H002",
				Price: 88,
			},
		}, map[int]*tesla.Modely{
			1 + 100*engine: {
				Name: "Y001",
				Mpg:  23.05,
			},
			2 + 200*engine: {
				Name: "Y002",
				Mpg:  33.60,
			},
		}
}
