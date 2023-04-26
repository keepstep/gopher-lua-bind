// Auto Generated !!!
// Type ModelY implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla/modely"
	lua "github.com/yuin/gopher-lua"
)

func Lua_ModelY_New(L *lua.LState) int {
	ins := &modely.ModelY{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("modely_ud"))
	L.Push(ud)
	return 1
}

func Lua_ModelY_ToUserData(L *lua.LState, ins *modely.ModelY) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("modely_ud"))
	return ud
}

func Lua_ModelY_LvToPtr(L *lua.LState, v lua.LValue) *modely.ModelY {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*modely.ModelY); ok {
			return ins
		} else {
			L.RaiseError("modely.ModelY expected")
			return nil
		}
	}
	return nil
}

func Lua_ModelY_UdToPtr(L *lua.LState, ud *lua.LUserData) *modely.ModelY {
	if ins, ok := ud.Value.(*modely.ModelY); ok {
		return ins
	} else {
		L.RaiseError("modely.ModelY expected")
		return nil
	}
}

func Lua_ModelY_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_ModelY_Check(L *lua.LState, n int) *modely.ModelY {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*modely.ModelY); ok {
		return v
	}
	L.ArgError(n, "modely expected")
	return nil
}

// funcs

// methods

//modely:Run(int) returns (int, error)
func Lua_ModelY_Run(L *lua.LState) int {
	ins := Lua_ModelY_Check(L, 1)
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

func Lua_ModelY_GetSet_Speed(L *lua.LState) int {
	ins := Lua_ModelY_Check(L, 1)
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

// Preload adds modely to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local modely = require("modely")
func PreloadModelY(L *lua.LState) {
	L.PreloadModule("modely", LoaderModelY)
}

// Loader is the module loader function.
func LoaderModelY(L *lua.LState) int {

	modely_ud := L.NewTypeMetatable("modely_ud")
	L.SetGlobal("modely_ud", modely_ud)
	L.SetField(modely_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		"Run": Lua_ModelY_Run,

		//field

		"Speed": Lua_ModelY_GetSet_Speed,

		//field slice

		//field map

	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_ModelY_New,
		//func

	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
