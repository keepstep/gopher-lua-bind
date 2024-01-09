package template

const TmpInit = `
// Auto Generated !!!
// Type implements for lua. 
package bind

import (
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

type TN interface {
	int8 | int16 | int32 | int64 | int | uint8 | uint16 | uint32 | uint64 | uint | float32 | float64
}

type TM interface {
	TN | bool | string
}

func Lua_SliceNumber_ToTable[T TN](L *lua.LState, m []T) *lua.LTable {
	tb := L.NewTable()
	for _, v := range m {
		vv := lua.LNumber(v)
		tb.Append(vv)
	}
	return tb
}

func Lua_SliceString_ToTable(L *lua.LState, m []string) *lua.LTable {
	tb := L.NewTable()
	for _, v := range m {
		vv := lua.LString(v)
		tb.Append(vv)
	}
	return tb
}

func Lua_SliceBool_ToTable(L *lua.LState, m []bool) *lua.LTable {
	tb := L.NewTable()
	for _, v := range m {
		vv := lua.LBool(v)
		tb.Append(vv)
	}
	return tb
}

func Lua_SliceError_ToTable(L *lua.LState, m []error) *lua.LTable {
	tb := L.NewTable()
	for _, v := range m {
		vv := lua.LString(v.Error())
		tb.Append(vv)
	}
	return tb
}

func Lua_Map_ToTable[K TM, V TM | any | map[K]V](L *lua.LState, m map[K]V) *lua.LTable {
	tb := L.NewTable()
	for k, v := range m {
		kkd := reflect.TypeOf(k).Kind()
		var klua lua.LValue = nil
		switch kkd {
		case reflect.Float32, reflect.Float64:
			klua = lua.LNumber(reflect.ValueOf(k).Float())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			klua = lua.LNumber(reflect.ValueOf(k).Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			klua = lua.LNumber(reflect.ValueOf(k).Uint())
		case reflect.String:
			klua = lua.LString(reflect.ValueOf(k).String())
		case reflect.Bool:
			klua = lua.LBool(reflect.ValueOf(k).Bool())
		default:
			return nil
		}

		vkd := reflect.TypeOf(v).Kind()
		var vlua lua.LValue = nil
		switch vkd {
		case reflect.Float32, reflect.Float64:
			vlua = lua.LNumber(reflect.ValueOf(v).Float())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vlua = lua.LNumber(reflect.ValueOf(v).Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			vlua = lua.LNumber(reflect.ValueOf(v).Uint())
		case reflect.String:
			vlua = lua.LString(reflect.ValueOf(v).String())
		case reflect.Bool:
			vlua = lua.LBool(reflect.ValueOf(v).Bool())
		case reflect.Map:
			vv := reflect.ValueOf(v).Interface()
			if sm, ok := vv.(map[K]V); ok {
				vlua = Lua_Map_ToTable(L, sm)
			}
		default:
			return nil
		}
		tb.RawSet(klua, vlua)
	}
	return tb
}

func Lua_Any_ToLValue(v any) lua.LValue {
	if v == nil {
		return lua.LNil
	}
	vkd := reflect.TypeOf(v).Kind()
	var vlua lua.LValue = nil
	switch vkd {
	case reflect.Float32, reflect.Float64:
		vlua = lua.LNumber(reflect.ValueOf(v).Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vlua = lua.LNumber(reflect.ValueOf(v).Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vlua = lua.LNumber(reflect.ValueOf(v).Uint())
	case reflect.String:
		vlua = lua.LString(reflect.ValueOf(v).String())
	case reflect.Bool:
		vlua = lua.LBool(reflect.ValueOf(v).Bool())
	default:
		return lua.LNil
	}
	return vlua
}

//any as golang basic type
//otherwise may cause err eg: excel SetCellValue
func Lua_LValueToAny(v lua.LValue) any {
	if v == nil {
		return nil
	}
	if v == lua.LNil {
		return nil
	}
	switch v.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return bool(lua.LVAsBool(v))
	case lua.LTNumber:
		return float64(lua.LVAsNumber(v))
	case lua.LTString:
		return string(lua.LVAsString(v))
	default:
		return nil
	}
}

//any as golang basic type
//otherwise may cause err eg: excel SetCellValue
func Lua_Any_Check(ls *lua.LState, n int) any {
	if n > ls.GetTop() {
		ls.ArgError(n, "value expected")
	}
	v := ls.Get(n)
	switch v.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return bool(lua.LVAsBool(v))
	case lua.LTNumber:
		return float64(lua.LVAsNumber(v))
	case lua.LTString:
		return string(lua.LVAsString(v))
	default:
		return nil
	}
}

func Lua_LValueToAny_Reflect(v lua.LValue, kd reflect.Kind) any {
	if v == nil {
		return nil
	}
	if v == lua.LNil {
		return nil
	}
	switch v.Type() {
	case lua.LTNil:
		return nil
	case lua.LTBool:
		return bool(lua.LVAsBool(v))
	case lua.LTNumber:
		if kd == reflect.Int {
			return int(lua.LVAsNumber(v))
		}
		if kd == reflect.Int8 {
			return int8(lua.LVAsNumber(v))
		}
		if kd == reflect.Int16 {
			return int16(lua.LVAsNumber(v))
		}
		if kd == reflect.Int32 {
			return int32(lua.LVAsNumber(v))
		}
		if kd == reflect.Int64 {
			return int64(lua.LVAsNumber(v))
		}
		if kd == reflect.Uint {
			return uint(lua.LVAsNumber(v))
		}
		if kd == reflect.Uint8 {
			return uint8(lua.LVAsNumber(v))
		}
		if kd == reflect.Uint16 {
			return uint16(lua.LVAsNumber(v))
		}
		if kd == reflect.Uint32 {
			return uint32(lua.LVAsNumber(v))
		}
		if kd == reflect.Uint64 {
			return uint64(lua.LVAsNumber(v))
		}
	case lua.LTString:
		return string(lua.LVAsString(v))
	default:
		return nil
	}
	return nil
}

func Lua_Slice_Check_JsonStr[T TM](L *lua.LState, n int) []T {
	m := []T{}
	s := L.CheckString(n)
	json.Unmarshal([]byte(s), &m)
	return m
}

func Lua_Map_Check_JsonStr[K TM, V TM | any | map[K]V](L *lua.LState, n int) map[K]V {
	m := map[K]V{}
	s := L.CheckString(n)
	json.Unmarshal([]byte(s), &m)
	return m
}

func Lua_Slice_Check[T TM](L *lua.LState, n int) []T {
	m := []T{}
	tb := L.CheckTable(n)
	var a T
	kd := reflect.TypeOf(a).Kind()
	tb.ForEach(func(k, v lua.LValue) {
		var vv any
		switch v.Type() {
		case lua.LTBool:
			fallthrough
		case lua.LTString:
			fallthrough
		case lua.LTNumber:
			vv = Lua_LValueToAny_Reflect(v, kd)
			if vv != nil {
				m = append(m, vv.(T))
			}
		}

	})
	return m
}

func Lua_Map_Check[K TM, V TM | any](L *lua.LState, n int) map[K]V {
	m := map[K]V{}
	tb := L.CheckTable(n)
	var a K
	var b K
	kkd := reflect.TypeOf(a).Kind()
	vkd := reflect.TypeOf(b).Kind()
	tb.ForEach(func(k, v lua.LValue) {
		var kk any
		switch k.Type() {
		case lua.LTBool:
			fallthrough
		case lua.LTString:
			fallthrough
		case lua.LTNumber:
			kk = Lua_LValueToAny_Reflect(k, kkd)
		}

		var vv any
		switch v.Type() {
		case lua.LTBool:
			fallthrough
		case lua.LTString:
			fallthrough
		case lua.LTNumber:
			vv = Lua_LValueToAny_Reflect(v, vkd)
		}
		if kk != nil && vv != nil {
			m[kk.(K)] = vv.(V)
		}
	})
	return m
}

func Preload(L *lua.LState) {
	{{ range $i,$o := .Objs }} 
	Preload{{ $o.Name }}(L)
	{{end }}

	{{ range $i,$o := .AllItf }} 
	{{ if eq $o.PkgPath "" }}
	{{ else }}
	PreloadItf{{ $o.Name }}(L)
	{{ end }}
	{{end }}
}
`
