// Auto Generated !!!
// Type BBB implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla/aaaa"
	lua "github.com/yuin/gopher-lua"
)

func Lua_BBB_New(L *lua.LState) int {
	ins := &aaaa.BBB{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("bbb_ud"))
	L.Push(ud)
	return 1
}

func Lua_BBB_ToUserData(L *lua.LState, ins *aaaa.BBB) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("bbb_ud"))
	return ud
}

func Lua_BBB_LvToPtr(L *lua.LState, v lua.LValue) *aaaa.BBB {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*aaaa.BBB); ok {
			return ins
		} else {
			L.RaiseError("aaaa.BBB expected")
			return nil
		}
	}
	return nil
}

func Lua_BBB_UdToPtr(L *lua.LState, ud *lua.LUserData) *aaaa.BBB {
	if ins, ok := ud.Value.(*aaaa.BBB); ok {
		return ins
	} else {
		L.RaiseError("aaaa.BBB expected")
		return nil
	}
}

func Lua_BBB_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_BBB_Check(L *lua.LState, n int) *aaaa.BBB {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*aaaa.BBB); ok {
		return v
	}
	L.ArgError(n, "bbb expected")
	return nil
}

// funcs

// methods

// fields

func Lua_BBB_GetSet_Tag(L *lua.LState) int {
	ins := Lua_BBB_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Tag = float64(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(ins.Tag))
	return 1
}

////////////////////////////////////////////////////////////////////

// Preload adds bbb to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local bbb = require("bbb")
func PreloadBBB(L *lua.LState) {
	L.PreloadModule("bbb", LoaderBBB)
}

// Loader is the module loader function.
func LoaderBBB(L *lua.LState) int {

	bbb_ud := L.NewTypeMetatable("bbb_ud")
	L.SetGlobal("bbb_ud", bbb_ud)
	L.SetField(bbb_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		//field

		"Tag": Lua_BBB_GetSet_Tag,

		//field slice

		//field map

	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_BBB_New,
		//func

	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
