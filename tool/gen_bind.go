package main

import (
	"fmt"

	"github.com/keepstep/gopher-lua-bind/car"
	"github.com/keepstep/gopher-lua-bind/car/byd"
	"github.com/keepstep/gopher-lua-bind/car/tesla"
	"github.com/keepstep/gopher-lua-bind/gen"
)

func main() {
	// objs := []gen.GenItem{
	// 	{Obj: &aaaa.AAA{}, Funcs: [][2]any{
	// 		{"AAACmp", aaaa.AAACmp},
	// 	}},
	// 	{Obj: &aaaa.BBB{}, Funcs: nil},
	// 	{Obj: &aaaa.CCC{}, Funcs: nil},
	// 	{Obj: &modely.ModelY{}, Funcs: nil},
	// 	{Obj: &modelx.ModelX{}, Funcs: nil},
	// 	{Obj: &modele.ModelE{}, Funcs: nil},
	// 	{Obj: &tesla.Tesla{}, Funcs: [][2]any{
	// 		{"TeslaCompare", tesla.TeslaCompare},
	// 		{"TeslaTest", tesla.TeslaTest},
	// 		{"TeslaGetAAA", tesla.TeslaGetAAA},
	// 		{"TeslaCallBack", tesla.TeslaCallBack},
	// 		{"TeslaCallInterface", tesla.TeslaCallInterface},
	// 		{"TeslaGetCmp", tesla.TeslaGetCmp},
	// 	}},
	// }
	objs2 := []gen.GenItem{
		{Obj: &car.Driver{}, Funcs: [][2]any{
			{"GetBrand", car.GetBrand},
			{"GetCars", car.GetCars},
		}},
		{Obj: &byd.Han{}, Funcs: nil},
		{Obj: &tesla.Modely{}, Funcs: nil},
	}
	err := gen.Gen(objs2, []string{}, "./", true)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
