// Auto Generated !!!
// Type Tesla implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla"
	lua "github.com/yuin/gopher-lua"

	"errors"
)

func Lua_Tesla_New(L *lua.LState) int {
	ins := &tesla.Tesla{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("tesla_ud"))
	L.Push(ud)
	return 1
}

func Lua_Tesla_ToUserData(L *lua.LState, ins *tesla.Tesla) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("tesla_ud"))
	return ud
}

func Lua_Tesla_LvToPtr(L *lua.LState, v lua.LValue) *tesla.Tesla {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*tesla.Tesla); ok {
			return ins
		} else {
			L.RaiseError("tesla.Tesla expected")
			return nil
		}
	}
	return nil
}

func Lua_Tesla_UdToPtr(L *lua.LState, ud *lua.LUserData) *tesla.Tesla {
	if ins, ok := ud.Value.(*tesla.Tesla); ok {
		return ins
	} else {
		L.RaiseError("tesla.Tesla expected")
		return nil
	}
}

func Lua_Tesla_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_Tesla_Check(L *lua.LState, n int) *tesla.Tesla {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*tesla.Tesla); ok {
		return v
	}
	L.ArgError(n, "tesla expected")
	return nil
}

// funcs

//tesla:TeslaCompare(string) returns (int)
func Lua_Tesla_TeslaCompare(L *lua.LState) int {

	p1 := L.CheckString(1)
	r1 := tesla.TeslaCompare(p1)

	L.Push(lua.LNumber(r1))

	return 1
}

//tesla:TeslaTest(int, float32, map[string]int) returns (int, float64, []string, error)
func Lua_Tesla_TeslaTest(L *lua.LState) int {

	p1 := L.CheckInt64(1)

	p2 := L.CheckNumber(2)

	p3ToMap := func(L *lua.LState, n int) map[string]int {
		m := map[string]int{}

		tb := L.CheckTable(n)
		tb.ForEach(func(k, v lua.LValue) {

			kk := lua.LVAsString(k)

			vv := int(lua.LVAsNumber(v))

			m[kk] = vv
		})
		return m
	}

	p3 := p3ToMap(L, 3)
	r1, r2, rt3, r4 := tesla.TeslaTest(int(p1), float32(p2), p3)

	L.Push(lua.LNumber(r1))

	L.Push(lua.LNumber(r2))

	rt3SliceToTable := func(L *lua.LState, m []string) *lua.LTable {

		tb := L.NewTable()
		for _, v := range m {

			vv := lua.LString(v)

			tb.Append(vv)
		}
		return tb
	}
	r3 := rt3SliceToTable(L, rt3)

	if r3 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(r3)
	}

	if r4 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r4.Error()))
	}

	return 4
}

//tesla:TeslaGetAAA(string) returns (AAA)
func Lua_Tesla_TeslaGetAAA(L *lua.LState) int {

	p1 := L.CheckString(1)
	r1 := tesla.TeslaGetAAA(p1)

	if r1 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(Lua_AAA_ToUserData(L, r1))
	}

	return 1
}

// methods

//tesla:Run(string, int) returns (int, error)
func Lua_Tesla_Run(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := L.CheckString(2)

	p2 := L.CheckInt64(3)
	r1, r2 := ins.Run(p1, int(p2))

	L.Push(lua.LNumber(r1))

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r2.Error()))
	}

	return 2
}

//tesla:RunE(ModelE) returns (int, error)
func Lua_Tesla_RunE(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := Lua_ModelE_Check(L, 2)
	r1, r2 := ins.RunE(p1)

	L.Push(lua.LNumber(r1))

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r2.Error()))
	}

	return 2
}

//tesla:RunX(ModelX) returns (int, error)
func Lua_Tesla_RunX(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := Lua_ModelX_Check(L, 2)
	r1, r2 := ins.RunX(p1)

	L.Push(lua.LNumber(r1))

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r2.Error()))
	}

	return 2
}

// fields

func Lua_Tesla_GetSet_Err(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Err = errors.New(L.CheckString(2))
		return 0
	}
	L.Push(Lua_Tesla_ErrorToLv(ins.Err))
	return 1
}

func Lua_Tesla_GetSet_Name(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Name = string(L.CheckString(2))
		return 0
	}
	L.Push(lua.LString(ins.Name))
	return 1
}

func Lua_Tesla_GetSet_Cars(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {

		p1ToSlice := func(L *lua.LState, n int) []string {
			m := []string{}

			tb := L.CheckTable(n)
			tb.ForEach(func(k, v lua.LValue) {

				vv := lua.LVAsString(v)

				m = append(m, vv)
			})
			return m
		}

		ins.Cars = p1ToSlice(L, 2)
		return 0
	}

	rt1SliceToTable := func(L *lua.LState, m []string) *lua.LTable {

		tb := L.NewTable()
		for _, v := range m {

			vv := lua.LString(v)

			tb.Append(vv)
		}
		return tb
	}
	r1 := rt1SliceToTable(L, ins.Cars)
	if r1 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(r1)
	}
	return 1
}

func Lua_Tesla_GetSet_Models(L *lua.LState) int {
	ins := Lua_Tesla_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {

		p1ToMap := func(L *lua.LState, n int) map[string]int {
			m := map[string]int{}

			tb := L.CheckTable(n)
			tb.ForEach(func(k, v lua.LValue) {

				kk := lua.LVAsString(k)

				vv := int(lua.LVAsNumber(v))

				m[kk] = vv
			})
			return m
		}
		ins.Models = p1ToMap(L, 2)
		return 0
	}

	r1MapToTable := func(L *lua.LState, m map[string]int) *lua.LTable {

		tb := L.NewTable()
		for k, v := range m {

			kk := lua.LString(k)

			vv := lua.LNumber(v)

			tb.RawSet(kk, vv)
		}
		return tb
	}
	r1 := r1MapToTable(L, ins.Models)
	if r1 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(r1)
	}
	return 1
}

////////////////////////////////////////////////////////////////////

// Preload adds tesla to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local tesla = require("tesla")
func PreloadTesla(L *lua.LState) {
	L.PreloadModule("tesla", LoaderTesla)
}

// Loader is the module loader function.
func LoaderTesla(L *lua.LState) int {

	tesla_ud := L.NewTypeMetatable("tesla_ud")
	L.SetGlobal("tesla_ud", tesla_ud)
	L.SetField(tesla_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		"Run": Lua_Tesla_Run,

		"RunE": Lua_Tesla_RunE,

		"RunX": Lua_Tesla_RunX,

		//field

		"Err": Lua_Tesla_GetSet_Err,

		"Name": Lua_Tesla_GetSet_Name,

		//field slice

		"Cars": Lua_Tesla_GetSet_Cars,

		//field map

		"Models": Lua_Tesla_GetSet_Models,
	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_Tesla_New,
		//func

		"TeslaCompare": Lua_Tesla_TeslaCompare,

		"TeslaTest": Lua_Tesla_TeslaTest,

		"TeslaGetAAA": Lua_Tesla_TeslaGetAAA,
	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
