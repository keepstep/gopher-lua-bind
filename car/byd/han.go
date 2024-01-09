package byd

import "fmt"

type Han struct {
	Name      string
	Seats     uint
	Mpg       float32
	Price     int
	Sale      bool
	CurrSpeed int
}

func (t *Han) SpeedUp(add int) int {
	t.CurrSpeed += add
	return t.CurrSpeed
}
func (t *Han) Park() bool {
	return true
}

func (t *Han) Info() string {
	return fmt.Sprintf("BydHan-%s-%d", t.Name, t.Price)
}
