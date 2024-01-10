package main

import (
	"fmt"

	"github.com/keepstep/gopher-lua-bind/gen"

	//your own files
	"github.com/keepstep/gopher-lua-bind/car"
	"github.com/keepstep/gopher-lua-bind/car/byd"
	"github.com/keepstep/gopher-lua-bind/car/tesla"
)

func main() {
	objs := []gen.GenItem{
		//objs with global functions
		{Obj: &car.Driver{}, Funcs: [][2]any{
			{"GetByCb", car.GetByCb},
			{"GetBrand", car.GetBrand},
			{"GetCars", car.GetCars},
		}},
		{Obj: &byd.Han{}, Funcs: nil},
		{Obj: &tesla.Modely{}, Funcs: nil},
	}
	//[param or field type] that defined in other packages will not be ignore
	allowPkgPath := []string{}
	//[outDir/bind] will be clean first then created, generated files will be placed in it
	outDir := "./"
	//[genLuaSnippet] will gen some lua file with incorrect syntax at outDir/lua
	genLuaSnippet := true
	err := gen.Gen(objs, allowPkgPath, outDir, genLuaSnippet)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
