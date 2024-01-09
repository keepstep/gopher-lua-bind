package car

import (
	"errors"
	"fmt"

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

func GetBrand(madein string, engine int, speed float32) (s []float64, o map[string]string) {
	return []float64{1.1, 1.2, 1.3, 1.4}, map[string]string{"byd": "C", "toyota": "J", "honda": "J", "bba": "G"}
}

func GetCars(madein string, engine int, speed float32) (s []*byd.Han, o map[int]*tesla.Modely) {
	return nil, nil
}
