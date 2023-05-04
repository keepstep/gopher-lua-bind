package template

const TmpBindInterface = `
{{ $ud_name := .UdName }}
{{ $lower_name := .LowerName }}
{{ $Name := .Name }}
// Auto Generated !!!
// Interface {{ .Name }} implements for lua. 
package bind

import (
	lua "github.com/yuin/gopher-lua"
	"{{ .PkgPath }}"
	{{ range $path,$i := .Import }}
	"{{ $path }}"
	{{ end }}
)

func Lua_{{ .Name }}_ToUserData(L *lua.LState, ins {{ .PkgName }}.{{ .Name }}) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = ins
	L.SetMetatable(ud, L.GetTypeMetatable("{{ $ud_name }}"))
	return ud
}

func Lua_{{ .Name }}_LvToPtr(L *lua.LState, v lua.LValue ) {{ .PkgName }}.{{ .Name }} {
	if ud, ok := v.(*lua.LUserData); ok {
		if ins, ok := ud.Value.({{ .PkgName }}.{{ .Name }}); ok {
			return ins
		} else {
			L.RaiseError("{{ .PkgName }}.{{ .Name }} expected")
			return nil
		}
	}
	return nil
}

func Lua_{{ .Name }}_UdToPtr(L *lua.LState, ud *lua.LUserData ) {{ .PkgName }}.{{ .Name }} {
	if ins, ok := ud.Value.({{ .PkgName }}.{{ .Name }}); ok {
		return ins
	} else {
		L.RaiseError("{{ .PkgName }}.{{ .Name }} expected")
		return nil
	}
}


func Lua_{{ .Name }}_ErrorToLv(err error) lua.LValue {
	if err != nil {
		return lua.LString(err.Error())
	}
	return lua.LNil
}


func Lua_{{ .Name}}_Check(L *lua.LState, n int) {{ .PkgName }}.{{ .Name }} {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.({{ .PkgName }}.{{ .Name }}); ok {
		return v
	}
	L.ArgError(n, "{{ .LowerName }} expected")
	return nil
}

// methods
{{ range $idx,$fun := .Methods }} 
//{{ $lower_name }}:{{- $fun.Name -}}({{- $fun.InType -}}) returns ( {{- $fun.OutType -}} )
func Lua_{{ $Name }}_{{- $fun.Name -}}(L *lua.LState) int {
	ins := Lua_{{ $Name }}_Check(L, 1)
	if ins == nil {
		return 0
	}
	{{ range $i,$t := .In }} 
		{{ if $t.IsMap }}
			{{ template "p_to_map" $t }}
		{{ else if $t.IsSlice }}
			{{ template "p_to_slice" $t }}
		{{ end }}
		{{ $fun.InDefine $i -}}
	{{ end }}
	{{ $fun.OutRetStr }} ins.{{- $fun.Name -}}( {{- $fun.InParam -}} )

	{{ range $i,$t := $fun.OutRetArr }}
		{{ if (index $t 2).IsMap }}
			{{ template "r_map_to_table" (index $t 2) }}
		{{ else if (index $t 2).IsSlice }}
			{{ template "r_slice_to_table" (index $t 2) }}
		{{ end }}
		{{ if $fun.OutCanNil $i }}
			if {{ index $t 0 }} == nil {
				L.Push(lua.LNil)
			}else{
				L.Push({{ index $t 1 }})
			}
		{{ else }}
			L.Push({{ index $t 1 }})
		{{ end }}
	{{ end }}

	return {{ $fun.OutLen }}
}
{{ end }}



////////////////////////////////////////////////////////////////////

// Preload adds {{ .LowerName }} to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local {{ .LowerName }} = require("{{ .LowerName }}")
func PreloadItf{{ .Name }}(L *lua.LState) {
	 LoaderItf{{ .Name }}(L)
}

// Loader is the module loader function.
func LoaderItf{{ .Name }}(L *lua.LState) int {

	{{ $ud_name }} := L.NewTypeMetatable("{{ $ud_name }}")
	L.SetGlobal("{{ $ud_name }}", {{ $ud_name }})
	L.SetField({{ $ud_name }}, "__index", L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
		//method
		{{ range $idx,$fun := .Methods }} 
			"{{- $fun.Name -}}" :  Lua_{{ $Name }}_{{- $fun.Name -}}, 
		{{ end }}
	}))
	
	return 0
}

{{- /*  mt := L.NewTypeMetatable("aaa_ud_mt") */ -}}
{{- /*  L.SetField(mt, "__call", L.NewFunction(Lua_AAA_New)) */ -}}
{{- /*  L.SetMetatable(t, mt) */ -}}


{{ define "p_to_map" }}
{{ $idx := .Index }}
{{ $key := .ElemKeyType }}
{{ $value := .ElemType }}
	{{ if $value.IsStruct }}
		p{{ $idx }}ToMap :=  func (L *lua.LState, n int) map[{{ $key.Name }}]*{{ $value.PkgName }}.{{ $value.Name }} {
			m := map[{{ $key.Name }}]*{{ $value.PkgName }}.{{ $value.Name }}{}
	{{ else }}
		p{{ $idx }}ToMap :=  func (L *lua.LState, n int) map[{{ $key.Name }}]{{ $value.Name }} {
			m := map[{{ $key.Name }}]{{ $value.Name }}{}
	{{ end }}
		tb := L.CheckTable(n)
		tb.ForEach(func(k, v lua.LValue) {
			{{ if $key.IsString }}
				kk := lua.LVAsString(k)
			{{ else if $key.IsNumber }}
				kk := {{ $key.Name }}(lua.LVAsNumber(k))
			{{ else if  $key.IsBool }}
				kk := lua.LVAsBool(k)
			{{ end }}

			{{ if $value.IsString }}
				vv := lua.LVAsString(v)
			{{ else if $value.IsNumber }}
				vv := {{ $value.Name }}(lua.LVAsNumber(v))
			{{ else if $value.IsBool }}
				vv := lua.LVAsBool(v)
			{{ else if $value.IsStruct }}
				vv := Lua_{{ $value.Name }}_LvToPtr(L,v)
			{{ end }}
			m[kk] = vv
		})
		return m
	}
{{ end }}

{{ define "p_to_slice" }}
{{ $idx := .Index }}
{{ $value := .ElemType }}
	{{ if $value.IsStruct }}
		p{{ $idx }}ToSlice :=  func (L *lua.LState, n int) []*{{ $value.PkgName }}.{{ $value.Name }} {
			m := []*{{ $value.PkgName }}.{{ $value.Name }}{}
	{{ else }}
		p{{ $idx }}ToSlice :=  func (L *lua.LState, n int) []{{ $value.Name }} {
			m := []{{ $value.Name }}{}
	{{ end }}
		tb := L.CheckTable(n)
		tb.ForEach(func(k, v lua.LValue) {
			{{ if $value.IsString }}
				vv := lua.LVAsString(v)
			{{ else if $value.IsNumber }}
				vv := {{ $value.Name }}(lua.LVAsNumber(v))
			{{ else if $value.IsBool }}
				vv := lua.LVAsBool(v)
			{{ else if $value.IsStruct }}
				vv := Lua_{{ $value.Name }}_LvToPtr(L,v)
			{{ end }}
			m = append(m,vv)
		})
		return m
	}
{{ end }}


{{ define "p_to_func" }}
{{ $idx := .Index }}
{{ $in := .In }}
{{ $out := .Out }}
	p{{ $idx }} :=  func (fp1 string) string {
		p := lua.P{
			Fn:   p2f,
			NRet: {{ len $out }},
		}
		lp1 := lua.LString(fp1)
		err := L.CallByParam(p, lp1)
		if err != nil {
			L.Error(lua.LString(err.Error()), 1)
		}
		ret := int(L.CheckNumber(1))
		return ret
	}
{{ end }}


{{ define "r_map_to_table" }}
{{ $idx := .Index }}
{{ $key := .ElemKeyType }}
{{ $value := .ElemType }}
	{{ if $value.IsStruct }}
		rt{{ $idx }}MapToTable :=  func (L *lua.LState, m map[{{ $key.Name }}]*{{ $value.PkgName }}.{{ $value.Name }}) *lua.LTable {
	{{ else }}
		rt{{ $idx }}MapToTable :=  func (L *lua.LState, m map[{{ $key.Name }}]{{ $value.Name }}) *lua.LTable {
	{{ end }}
		tb := L.NewTable()
		for k,v := range m {
			{{ if $key.IsString }}
				kk := lua.LString(k)
			{{ else if $key.IsNumber }}
				kk := lua.LNumber(k)
			{{ else if  $key.IsBool }}
				kk := lua.LBool(k)
			{{ end }}

			{{ if $value.IsString }}
				vv := lua.LString(v)
			{{ else if $value.IsNumber }}
				vv := lua.LNumber(v)
			{{ else if $value.IsBool }}
				vv := lua.LBool(v)
			{{ else if $value.IsStruct }}
				vv := Lua_{{ $value.Name }}_ToUserData(L,v)	
			{{ end }}
			tb.RawSet(kk,vv)
		}
		return tb
	}
	r{{ $idx }} := rt{{ $idx }}MapToTable(L,rt{{ $idx }})
{{ end }}

{{ define "r_slice_to_table" }}
{{ $idx := .Index }}
{{ $value := .ElemType }}
	{{ if $value.IsStruct }}
		rt{{ $idx }}SliceToTable :=  func (L *lua.LState, m []*{{ $value.PkgName }}.{{ $value.Name }}) *lua.LTable {
	{{ else }}
		rt{{ $idx }}SliceToTable :=  func (L *lua.LState, m []{{ $value.Name }}) *lua.LTable {
	{{ end }}
		tb := L.NewTable()
		for _,v := range m {
			{{ if $value.IsString }}
				vv := lua.LString(v)
			{{ else if $value.IsNumber }}
				vv := lua.LNumber(v)
			{{ else if $value.IsBool }}
				vv := lua.LBool(v)
			{{ else if $value.IsStruct }}
				vv := Lua_{{ $value.Name }}_ToUserData(L,v)	
			{{ end }}
			tb.Append(vv)
		}
		return tb
	}
	r{{ $idx }} := rt{{ $idx }}SliceToTable(L,rt{{ $idx }})
{{ end }}


{{ define "field_func_slice" }}
{{ $Name := index . 0 }}
{{ $fn := index . 1 }}
{{ $value := (index . 2).ElemType }}
func Lua_{{ $Name }}_GetSet_{{ $fn }}(L *lua.LState) int {
	ins := Lua_{{ $Name }}_Check(L, 1)
	if ins == nil {
		return 0
	}
    if L.GetTop() == 2 {
		{{ if $value.IsStruct }}
			p1ToSlice :=  func (L *lua.LState, n int) []*{{ $value.PkgName }}.{{ $value.Name }} {
			m := []*{{ $value.PkgName }}.{{ $value.Name }}{}
		{{ else }}
			p1ToSlice :=  func (L *lua.LState, n int) []{{ $value.Name }} {
				m := []{{ $value.Name }}{}
		{{ end }}
			tb := L.CheckTable(n)
			tb.ForEach(func(k, v lua.LValue) {
				{{ if $value.IsString }}
					vv := lua.LVAsString(v)
				{{ else if $value.IsNumber }}
					vv := {{ $value.Name }}(lua.LVAsNumber(v))
				{{ else if $value.IsBool }}
					vv := lua.LVAsBool(v)
				{{ else if $value.IsStruct }}
					vv := Lua_{{ $value.Name }}_LvToPtr(L,v)
				{{ end }}
				m = append(m,vv)
			})
			return m
		}

        ins.{{ $fn }} = p1ToSlice(L,2)
        return 0
    }
	{{ if $value.IsStruct }}
		rt1SliceToTable :=  func (L *lua.LState, m []*{{ $value.PkgName }}.{{ $value.Name }}) *lua.LTable {
	{{ else }}
		rt1SliceToTable :=  func (L *lua.LState, m []{{ $value.Name }}) *lua.LTable {
	{{ end }}
		tb := L.NewTable()
		for _,v := range m {
			{{ if $value.IsString }}
				vv := lua.LString(v)
			{{ else if $value.IsNumber }}
				vv := lua.LNumber(v)
			{{ else if $value.IsBool }}
				vv := lua.LBool(v)
			{{ else if $value.IsStruct }}
				vv := Lua_{{ $value.Name }}_ToUserData(L,v)	
			{{ end }}
			tb.Append(vv)
		}
		return tb
	}
	r1 := rt1SliceToTable(L,ins.{{ $fn }})
	if r1 == nil {
		L.Push(lua.LNil)
	}else{
		L.Push(r1)
	}
    return 1
}	
{{ end }}

{{ define "field_func_map" }}
{{ $Name := index . 0 }}
{{ $fn := index . 1 }}
{{ $key := (index . 2).ElemKeyType }}
{{ $value := (index . 2).ElemType }}
func Lua_{{ $Name }}_GetSet_{{ $fn }}(L *lua.LState) int {
	ins := Lua_{{ $Name }}_Check(L, 1)
	if ins == nil {
		return 0
	}
    if L.GetTop() == 2 {
		{{ if $value.IsStruct }}
			p1ToMap :=  func (L *lua.LState, n int) map[{{ $key.Name }}]*{{ $value.PkgName }}.{{ $value.Name }} {
			m := map[{{ $key.Name }}]*{{ $value.PkgName }}.{{ $value.Name }}{}
		{{ else }}
			p1ToMap :=  func (L *lua.LState, n int) map[{{ $key.Name }}]{{ $value.Name }} {
				m := map[{{ $key.Name }}]{{ $value.Name }}{}
		{{ end }}
			tb := L.CheckTable(n)
			tb.ForEach(func(k, v lua.LValue) {
				{{ if $key.IsString }}
					kk := lua.LVAsString(k)
				{{ else if $key.IsNumber }}
					kk := {{ $key.Name }}(lua.LVAsNumber(k))
				{{ else if  $key.IsBool }}
					kk := lua.LVAsBool(k)
				{{ end }}

				{{ if $value.IsString }}
					vv := lua.LVAsString(v)
				{{ else if $value.IsNumber }}
					vv := {{ $value.Name }}(lua.LVAsNumber(v))
				{{ else if $value.IsBool }}
					vv := lua.LVAsBool(v)
				{{ else if $value.IsStruct }}
					vv := Lua_{{ $value.Name }}_LvToPtr(L,v)
				{{ end }}
				m[kk] = vv
			})
			return m
		}
        ins.{{ $fn }} = p1ToMap(L,2)
        return 0
    }

	{{ if $value.IsStruct }}
		r1MapToTable :=  func (L *lua.LState, m map[{{ $key.Name }}]*{{ $value.PkgName }}.{{ $value.Name }}) *lua.LTable {
	{{ else }}
		r1MapToTable :=  func (L *lua.LState, m map[{{ $key.Name }}]{{ $value.Name }}) *lua.LTable {
	{{ end }}
		tb := L.NewTable()
		for k,v := range m {
			{{ if $key.IsString }}
				kk := lua.LString(k)
			{{ else if $key.IsNumber }}
				kk := lua.LNumber(k)
			{{ else if  $key.IsBool }}
				kk := lua.LBool(k)
			{{ end }}

			{{ if $value.IsString }}
				vv := lua.LString(v)
			{{ else if $value.IsNumber }}
				vv := lua.LNumber(v)
			{{ else if $value.IsBool }}
				vv := lua.LBool(v)
			{{ else if $value.IsStruct }}
				vv := Lua_{{ $value.Name }}_ToUserData(L,v)	
			{{ end }}
			tb.RawSet(kk,vv)
		}
		return tb
	}
	r1 := r1MapToTable(L,ins.{{ $fn }})
	if r1 == nil {
		L.Push(lua.LNil)
	}else{
		L.Push(r1)
	}
    return 1
}	
{{ end }}


{{ define "field_func_func" }}
{{ $Name := index . 0 }}
{{ $fn := index . 1 }}
{{ $type := (index . 2)}}
func Lua_{{ $Name }}_Set_{{ $fn }}(L *lua.LState) int {
	ins := Lua_{{ $Name }}_Check(L, 1)
	if ins == nil {
		return 0
	}
    if L.GetTop() == 2 {
		p1f := L.CheckFunction(2)
		p1 := {{ $type.FuncDefine "p1f" }}
        ins.{{ $fn }} = p1
        return 0
    }
	L.Push(lua.LNil)
    return 1
}	
{{ end }}

`
