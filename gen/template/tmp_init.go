package template

const TmpInit = `
// Auto Generated !!!
// Type implements for lua. 
package bind

import (
	lua "github.com/yuin/gopher-lua"
)

func Preload(L *lua.LState) {
	{{ range $i,$o := .Objs }} 
	Preload{{ $o.Name }}(L)
	{{end }}
}
`
