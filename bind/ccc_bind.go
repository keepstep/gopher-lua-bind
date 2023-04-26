// Auto Generated !!!
// Type CCC implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla/aaaa"
	lua "github.com/yuin/gopher-lua"
)

func Lua_CCC_New(L *lua.LState) int {
	ins := &aaaa.CCC{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("ccc_ud"))
	L.Push(ud)
	return 1
}

func Lua_CCC_ToUserData(L *lua.LState, ins *aaaa.CCC) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("ccc_ud"))
	return ud
}

func Lua_CCC_LvToPtr(L *lua.LState, v lua.LValue) *aaaa.CCC {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*aaaa.CCC); ok {
			return ins
		} else {
			L.RaiseError("aaaa.CCC expected")
			return nil
		}
	}
	return nil
}

func Lua_CCC_UdToPtr(L *lua.LState, ud *lua.LUserData) *aaaa.CCC {
	if ins, ok := ud.Value.(*aaaa.CCC); ok {
		return ins
	} else {
		L.RaiseError("aaaa.CCC expected")
		return nil
	}
}

func Lua_CCC_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_CCC_Check(L *lua.LState, n int) *aaaa.CCC {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*aaaa.CCC); ok {
		return v
	}
	L.ArgError(n, "ccc expected")
	return nil
}

// funcs

// methods

// fields

func Lua_CCC_GetSet_Code(L *lua.LState) int {
	ins := Lua_CCC_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Code = int64(L.CheckInt64(2))
		return 0
	}
	L.Push(lua.LNumber(ins.Code))
	return 1
}

////////////////////////////////////////////////////////////////////

// Preload adds ccc to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local ccc = require("ccc")
func PreloadCCC(L *lua.LState) {
	L.PreloadModule("ccc", LoaderCCC)
}

// Loader is the module loader function.
func LoaderCCC(L *lua.LState) int {

	ccc_ud := L.NewTypeMetatable("ccc_ud")
	L.SetGlobal("ccc_ud", ccc_ud)
	L.SetField(ccc_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		//field

		"Code": Lua_CCC_GetSet_Code,

		//field slice

		//field map

	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_CCC_New,
		//func

	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
