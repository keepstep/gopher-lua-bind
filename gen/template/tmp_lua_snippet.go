package template

const TmpLuaSnippet = `-- Auto Generated !!!
-- Type {{ .Name }}'s lua snippet with incorrect syntax. 
{{ $ud_name := .UdName }}
{{ $lower_name := .LowerName }}
{{ $Name := .Name }}

local {{ $lower_name }} = require("{{- $lower_name -}}")

local obj = {{ $lower_name -}}.new()

-- fields
{{ range $fn,$ff := .FieldsBind }}
local f =  obj:{{ $fn }}()
obj:{{ $fn }}({{ index $ff 2 }})
{{ end }}

-- field struct ptr
{{ range $fn,$ff := .FieldsBindStructPtr }}
local f =  obj:{{ $fn }}()
obj:{{ $fn }}({{ index $ff 2 }})
{{ end }}

-- field interface
{{ range $fn,$ff := .FieldsBindInterface }}
local f =  obj:{{ $fn }}()
obj:{{ $fn }}({{ (index $ff 2).RefType }})
{{ end }}

-- field slice
{{ range $i,$ff := .FieldsBindSlice }}
{{ $name := (index $ff 2).ElemType.Name }}
local f =  obj:{{ index $ff 1 }}()
obj:{{ index $ff 1 }}([]{{ $name }})
{{ end }}

-- field map
{{ range $i,$ff := .FieldsBindMap }}
{{ $keyName := (index $ff 2).ElemKeyType.Name }}
{{ $valueName := (index $ff 2).ElemType.Name }}
local f =  obj:{{ index $ff 1 }}()
obj:{{ index $ff 1 }}(map[{{ $keyName }}]{{ $valueName }})
{{ end }}

-- field func
{{ range $i,$ff := .FieldsBindFunc }}
local f
obj:{{ index $ff 1  }}( {{ (index $ff 2).RefType }} )
f = obj:{{ index $ff 1 }}()
f()
{{ end }}

-- methods
{{ range $idx,$fun := .Methods }}
local {{ $fun.OutType2 }} = obj:{{- $fun.Name -}}({{- $fun.InType2 -}})
{{ end }}

-- funcs
{{ range $idx,$fun := .Funcs }}
local {{ $fun.OutType2 }} = {{ $lower_name }}.{{- $fun.Name -}}({{- $fun.InType2 -}})
{{ end }}
`
