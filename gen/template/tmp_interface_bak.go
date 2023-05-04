package template

const TmpInterface = `
// Auto Generated !!!
// Type interface implements for lua. 
package bind

import (
	lua "github.com/yuin/gopher-lua"
	{{ range $path,$i := .InterfacePkgPath }}
	"{{ $path }}"
	{{ end }}
)
{{ range $path,$t := .AllInterface }}
{{ if eq $t.PkgName  "" }}
func Lua_{{ $t.Name }}_Interface_Check(L *lua.LState, n int) {{ $t.Name }} {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.({{ $t.Name }}); ok {
		return v
	}
	L.ArgError(n, "{{ $t.Name }} expected")
	return nil
}
{{ else }}
func Lua_{{ $t.Name }}_Interface_Check(L *lua.LState, n int) {{ $t.PkgName }}.{{ $t.Name }} {
	ud := L.CheckUserData(n)
	if v, ok := ud.Value.({{ $t.PkgName }}.{{ $t.Name }}); ok {
		return v
	}
	L.ArgError(n, "{{ $t.PkgName }}.{{ $t.Name }} expected")
	return nil
}
{{ end }}
{{ end }}


`
