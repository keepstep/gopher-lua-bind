// Auto Generated !!!
// Type ModelX implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla/modelx"
	lua "github.com/yuin/gopher-lua"
)

func Lua_ModelX_New(L *lua.LState) int {
	ins := &modelx.ModelX{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("modelx_ud"))
	L.Push(ud)
	return 1
}

func Lua_ModelX_ToUserData(L *lua.LState, ins *modelx.ModelX) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("modelx_ud"))
	return ud
}

func Lua_ModelX_LvToPtr(L *lua.LState, v lua.LValue) *modelx.ModelX {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*modelx.ModelX); ok {
			return ins
		} else {
			L.RaiseError("modelx.ModelX expected")
			return nil
		}
	}
	return nil
}

func Lua_ModelX_UdToPtr(L *lua.LState, ud *lua.LUserData) *modelx.ModelX {
	if ins, ok := ud.Value.(*modelx.ModelX); ok {
		return ins
	} else {
		L.RaiseError("modelx.ModelX expected")
		return nil
	}
}

func Lua_ModelX_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_ModelX_Check(L *lua.LState, n int) *modelx.ModelX {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*modelx.ModelX); ok {
		return v
	}
	L.ArgError(n, "modelx expected")
	return nil
}

// funcs

// methods

//modelx:Run(int) returns (int, error)
func Lua_ModelX_Run(L *lua.LState) int {
	ins := Lua_ModelX_Check(L, 1)
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

func Lua_ModelX_GetSet_Speed(L *lua.LState) int {
	ins := Lua_ModelX_Check(L, 1)
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

// Preload adds modelx to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local modelx = require("modelx")
func PreloadModelX(L *lua.LState) {
	L.PreloadModule("modelx", LoaderModelX)
}

// Loader is the module loader function.
func LoaderModelX(L *lua.LState) int {

	modelx_ud := L.NewTypeMetatable("modelx_ud")
	L.SetGlobal("modelx_ud", modelx_ud)
	L.SetField(modelx_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		"Run": Lua_ModelX_Run,

		//field

		"Speed": Lua_ModelX_GetSet_Speed,

		//field slice

		//field map

	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_ModelX_New,
		//func

	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
