# gopher-lua-bind
[![GoDoc](https://godoc.org/github.com/keepstep/gopher-lua-bind?status.svg)](https://godoc.org/github.com/keepstep/gopher-lua-bind)


It is a simple bind-generation for [gopher-lua](https://github.com/yuin/gopher-lua) with limited features.  

## No Support Bind and 
```
1. Chan
2. *int *float ... as struct field
3. *int,*float,*string in function params
4. struct, etc. as callback func input params
```

## Installation

```
go get github.com/keepstep/gopher-lua-bind
```
## Test

```bash
# generate bind
sh gen_bind.sh
# run
sh run.sh
```

## Usage

Generate code 

```golang
// tool/gen_bind.go
package main

import (
        "fmt"

        "github.com/keepstep/gopher-lua-bind/gen"
        //your own files
        "github.com/keepstep/gopher-lua-bind/tesla"
        "github.com/keepstep/gopher-lua-bind/tesla/aaaa"
        "github.com/keepstep/gopher-lua-bind/tesla/modele"
        "github.com/keepstep/gopher-lua-bind/tesla/modelx"
        "github.com/keepstep/gopher-lua-bind/tesla/modely"
)

func main() {
        //objs with global functions
        objs := []gen.GenItem{
                {Obj: &aaaa.AAA{}, Funcs: nil},
                {Obj: &aaaa.BBB{}, Funcs: nil},
                {Obj: &aaaa.CCC{}, Funcs: nil},
                {Obj: &modely.ModelY{}, Funcs: nil},
                {Obj: &modelx.ModelX{}, Funcs: nil},
                {Obj: &modele.ModelE{}, Funcs: nil},
                {Obj: &tesla.Tesla{}, Funcs: [][2]any{
                        {"TeslaCompare", tesla.TeslaCompare},
                        {"TeslaTest", tesla.TeslaTest},
                        {"TeslaGetAAA", tesla.TeslaGetAAA},
                }},
        }
        //【param or field type】 that defined in other packages will not be ignore
        allowPkgPath := []string{}
        //【outDir/bind】 will be clean first then created, generated files will be placed in it
        outDir := "./"
        err := gen.Gen(objs,allowPkgPath , outDir)
        if err != nil {
                fmt.Printf("Error: %v\n", err)
        }
}


```


Run generate and format
```bash
rm -f ./gen_bind
go build -o gen_bind ./tool/gen_bind.go
./gen_bind
gofmt -l -w ./bind
```

Basic usage of bind.Preload in your code 

```golang
// main.go
package main

import (
	"flag"
	"log"
	"os"

	"github.com/keepstep/gopher-lua-bind/bind"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	flag.Parse()
	state := lua.NewState()
	defer state.Close()
	bind.Preload(state)
	state.DoFile("lua/test_bind.lua")
}
```


## License

MIT


