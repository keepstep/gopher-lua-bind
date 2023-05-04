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

	{{ range $i,$o := .AllItf }} 
	{{ if eq $o.PkgPath "" }}
	{{ else }}
	PreloadItf{{ $o.Name }}(L)
	{{ end }}
	{{end }}
}
`
