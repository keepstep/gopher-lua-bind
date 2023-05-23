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

func Lua_Map_ToTable[K TM, V TM](L *lua.LState, m map[K]V) *lua.LTable {
	tb := L.NewTable()
	for k, v := range m {
		kkd := reflect.TypeOf(k).Kind()
		vkd := reflect.TypeOf(v).Kind()
		var klua lua.LValue = nil
		var vlua lua.LValue = nil
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
		switch vkd {
		case reflect.Float32, reflect.Float64:
			vlua = lua.LNumber(reflect.ValueOf(k).Float())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vlua = lua.LNumber(reflect.ValueOf(k).Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			vlua = lua.LNumber(reflect.ValueOf(k).Uint())
		case reflect.String:
			vlua = lua.LString(reflect.ValueOf(k).String())
		case reflect.Bool:
			vlua = lua.LBool(reflect.ValueOf(k).Bool())
		default:
			return nil
		}
		tb.RawSet(klua, vlua)
	}
	return tb
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
