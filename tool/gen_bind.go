package main

import (
	"fmt"

	"github.com/keepstep/gopher-lua-bind/gen"
	"github.com/keepstep/gopher-lua-bind/tesla"
	"github.com/keepstep/gopher-lua-bind/tesla/aaaa"
	"github.com/keepstep/gopher-lua-bind/tesla/modele"
	"github.com/keepstep/gopher-lua-bind/tesla/modelx"
	"github.com/keepstep/gopher-lua-bind/tesla/modely"
)

func main() {
	objs := []gen.GenItem{
		{Obj: &aaaa.AAA{}, Funcs: nil},
		{Obj: &aaaa.BBB{}, Funcs: nil},
		{Obj: &aaaa.CCC{}, Funcs: nil},
		{Obj: &modely.ModelY{}, Funcs: nil},
		{Obj: &modelx.ModelX{}, Funcs: nil},
		{Obj: &modele.ModelE{}, Funcs: nil},
		{Obj: &tesla.Tesla{}, Funcs: [][2]any{
			{"TeslaCompare", tesla.TeslaCompare},
			{"TeslaTest", tesla.TeslaTest},
			{"TeslaGetAAA", tesla.TeslaGetAAA},
		}},
	}
	err := gen.Gen(objs, []string{}, "./")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
