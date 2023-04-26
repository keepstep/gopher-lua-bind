// Auto Generated !!!
// Type ModelE implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla/modele"
	lua "github.com/yuin/gopher-lua"
)

func Lua_ModelE_New(L *lua.LState) int {
	ins := &modele.ModelE{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("modele_ud"))
	L.Push(ud)
	return 1
}

func Lua_ModelE_ToUserData(L *lua.LState, ins *modele.ModelE) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("modele_ud"))
	return ud
}

func Lua_ModelE_LvToPtr(L *lua.LState, v lua.LValue) *modele.ModelE {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*modele.ModelE); ok {
			return ins
		} else {
			L.RaiseError("modele.ModelE expected")
			return nil
		}
	}
	return nil
}

func Lua_ModelE_UdToPtr(L *lua.LState, ud *lua.LUserData) *modele.ModelE {
	if ins, ok := ud.Value.(*modele.ModelE); ok {
		return ins
	} else {
		L.RaiseError("modele.ModelE expected")
		return nil
	}
}

func Lua_ModelE_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_ModelE_Check(L *lua.LState, n int) *modele.ModelE {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*modele.ModelE); ok {
		return v
	}
	L.ArgError(n, "modele expected")
	return nil
}

// funcs

// methods

//modele:Run(int) returns (int, error)
func Lua_ModelE_Run(L *lua.LState) int {
	ins := Lua_ModelE_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := L.CheckInt64(2)
	r1, r2 := ins.Run(int(p1))

	L.Push(lua.LNumber(r1))

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r2.Error()))
	}

	return 2
}

// fields

func Lua_ModelE_GetSet_Speed(L *lua.LState) int {
	ins := Lua_ModelE_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Speed = int(L.CheckInt64(2))
		return 0
	}
	L.Push(lua.LNumber(ins.Speed))
	return 1
}

////////////////////////////////////////////////////////////////////

// Preload adds modele to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local modele = require("modele")
func PreloadModelE(L *lua.LState) {
	L.PreloadModule("modele", LoaderModelE)
}

// Loader is the module loader function.
func LoaderModelE(L *lua.LState) int {

	modele_ud := L.NewTypeMetatable("modele_ud")
	L.SetGlobal("modele_ud", modele_ud)
	L.SetField(modele_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		"Run": Lua_ModelE_Run,

		//field

		"Speed": Lua_ModelE_GetSet_Speed,

		//field slice

		//field map

	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_ModelE_New,
		//func

	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
