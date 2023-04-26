// Auto Generated !!!
// Type AAA implements for lua.
package bind

import (
	"github.com/keepstep/gopher-lua-bind/tesla/aaaa"
	lua "github.com/yuin/gopher-lua"
)

func Lua_AAA_New(L *lua.LState) int {
	ins := &aaaa.AAA{}
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("aaa_ud"))
	L.Push(ud)
	return 1
}

func Lua_AAA_ToUserData(L *lua.LState, ins *aaaa.AAA) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("aaa_ud"))
	return ud
}

func Lua_AAA_LvToPtr(L *lua.LState, v lua.LValue) *aaaa.AAA {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.(*aaaa.AAA); ok {
			return ins
		} else {
			L.RaiseError("aaaa.AAA expected")
			return nil
		}
	}
	return nil
}

func Lua_AAA_UdToPtr(L *lua.LState, ud *lua.LUserData) *aaaa.AAA {
	if ins, ok := ud.Value.(*aaaa.AAA); ok {
		return ins
	} else {
		L.RaiseError("aaaa.AAA expected")
		return nil
	}
}

func Lua_AAA_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}

func Lua_AAA_Check(L *lua.LState, n int) *aaaa.AAA {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.(*aaaa.AAA); ok {
		return v
	}
	L.ArgError(n, "aaa expected")
	return nil
}

// funcs

// methods

//aaa:Do(int, BBB, CCC) returns (string, AAA, BBB, error)
func Lua_AAA_Do(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := L.CheckInt64(2)

	p2 := Lua_BBB_Check(L, 3)

	p3 := Lua_CCC_Check(L, 4)
	r1, r2, r3, r4 := ins.Do(int(p1), p2, p3)

	L.Push(lua.LString(r1))

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(Lua_AAA_ToUserData(L, r2))
	}

	if r3 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(Lua_BBB_ToUserData(L, r3))
	}

	if r4 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r4.Error()))
	}

	return 4
}

//aaa:DoAdd(int, float32) returns (int, error)
func Lua_AAA_DoAdd(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := L.CheckInt64(2)

	p2 := L.CheckNumber(3)
	r1, r2 := ins.DoAdd(int(p1), float32(p2))

	L.Push(lua.LNumber(r1))

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r2.Error()))
	}

	return 2
}

//aaa:DoMap(int, float32, map[string]int) returns (int, float64, map[string]string, error)
func Lua_AAA_DoMap(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1 := L.CheckInt64(2)

	p2 := L.CheckNumber(3)

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

	p3 := p3ToMap(L, 4)
	r1, r2, rt3, r4 := ins.DoMap(int(p1), float32(p2), p3)

	L.Push(lua.LNumber(r1))

	L.Push(lua.LNumber(r2))

	rt3MapToTable := func(L *lua.LState, m map[string]string) *lua.LTable {

		tb := L.NewTable()
		for k, v := range m {

			kk := lua.LString(k)

			vv := lua.LString(v)

			tb.RawSet(kk, vv)
		}
		return tb
	}
	r3 := rt3MapToTable(L, rt3)

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

//aaa:DoSlice([]int) returns ([]string, error)
func Lua_AAA_DoSlice(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
	if ins == nil {
		return 0
	}

	p1ToSlice := func(L *lua.LState, n int) []int {
		m := []int{}

		tb := L.CheckTable(n)
		tb.ForEach(func(k, v lua.LValue) {

			vv := int(lua.LVAsNumber(v))

			m = append(m, vv)
		})
		return m
	}

	p1 := p1ToSlice(L, 2)
	rt1, r2 := ins.DoSlice(p1)

	rt1SliceToTable := func(L *lua.LState, m []string) *lua.LTable {

		tb := L.NewTable()
		for _, v := range m {

			vv := lua.LString(v)

			tb.Append(vv)
		}
		return tb
	}
	r1 := rt1SliceToTable(L, rt1)

	if r1 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(r1)
	}

	if r2 == nil {
		L.Push(lua.LNil)
	} else {
		L.Push(lua.LString(r2.Error()))
	}

	return 2
}

// fields

func Lua_AAA_GetSet_Age(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Age = int(L.CheckInt64(2))
		return 0
	}
	L.Push(lua.LNumber(ins.Age))
	return 1
}

func Lua_AAA_GetSet_Name(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
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

func Lua_AAA_GetSet_Weight(L *lua.LState) int {
	ins := Lua_AAA_Check(L, 1)
	if ins == nil {
		return 0
	}
	if L.GetTop() == 2 {
		ins.Weight = float32(L.CheckNumber(2))
		return 0
	}
	L.Push(lua.LNumber(ins.Weight))
	return 1
}

////////////////////////////////////////////////////////////////////

// Preload adds aaa to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local aaa = require("aaa")
func PreloadAAA(L *lua.LState) {
	L.PreloadModule("aaa", LoaderAAA)
}

// Loader is the module loader function.
func LoaderAAA(L *lua.LState) int {

	aaa_ud := L.NewTypeMetatable("aaa_ud")
	L.SetGlobal("aaa_ud", aaa_ud)
	L.SetField(aaa_ud, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method

		"Do": Lua_AAA_Do,

		"DoAdd": Lua_AAA_DoAdd,

		"DoMap": Lua_AAA_DoMap,

		"DoSlice": Lua_AAA_DoSlice,

		//field

		"Age": Lua_AAA_GetSet_Age,

		"Name": Lua_AAA_GetSet_Name,

		"Weight": Lua_AAA_GetSet_Weight,

		//field slice

		//field map

	}))

	var api = map[string]lua.LGFunction{
		"new": Lua_AAA_New,
		//func

	}

	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}
