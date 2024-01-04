package gen

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	tmp "github.com/keepstep/gopher-lua-bind/gen/template"
)

type GenItem struct {
	Obj   any
	Funcs [][2]any
}

// func
// allowPkgPath
func Gen(items []GenItem, allowPkgPath []string, outDir string, genLuaSnippet bool) error {

	b := NewBindData(items, allowPkgPath)
	err := b.Load()
	if err != nil {
		panic(err)
	}
	bs, _ := json.MarshalIndent(b.AllItf, "", "    ")
	fmt.Println(string(bs))

	dirName := ""
	if outDir == "" {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		dirName = exPath + "/bind"
	} else {
		dirName = filepath.Dir(outDir) + "/bind"
	}
	fmt.Printf("dirName = %s \n", dirName)
	os.RemoveAll(dirName)
	err = os.Mkdir(dirName, 0666)
	if err != nil {
		panic(err)
	}

	luaDir := dirName + "/lua"
	if genLuaSnippet {
		err = os.Mkdir(luaDir, 0666)
		if err != nil {
			panic(err)
		}
	}

	for _, obj := range b.AllObj {
		bindFile := fmt.Sprintf("%s/struct_%s.go", dirName, obj.LowerName())
		t, err := template.New("bind").Parse(tmp.TmpStruct)
		if err != nil {
			return err
		}
		lf, err := os.Create(bindFile)
		if err != nil {
			panic(err)
		}
		defer func() {
			lf.Close()
		}()
		err = t.Execute(lf, obj)
		if err != nil {
			return err
		}
	}
	if genLuaSnippet {
		for _, obj := range b.AllObj {
			bindFile := fmt.Sprintf("%s/struct_%s.lua", luaDir, obj.LowerName())
			t, err := template.New("lua_snippet").Parse(tmp.TmpLuaSnippet)
			if err != nil {
				return err
			}
			lf, err := os.Create(bindFile)
			if err != nil {
				panic(err)
			}
			defer func() {
				lf.Close()
			}()
			err = t.Execute(lf, obj)
			if err != nil {
				return err
			}
		}
	}

	for _, obj := range b.AllItf {
		if obj.IsError {
			continue
		}
		bindFile := fmt.Sprintf("%s/interface_%s.go", dirName, obj.LowerName())
		t, err := template.New("itf").Parse(tmp.TmpBindInterface)
		if err != nil {
			return err
		}
		lf, err := os.Create(bindFile)
		if err != nil {
			panic(err)
		}
		defer func() {
			lf.Close()
		}()
		err = t.Execute(lf, obj)
		if err != nil {
			return err
		}
	}

	initFile := fmt.Sprintf("%s/init.go", dirName)
	t, err := template.New("init").Parse(tmp.TmpInit)
	if err != nil {
		return err
	}
	lf, err := os.Create(initFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		lf.Close()
	}()
	err = t.Execute(lf, b)
	if err != nil {
		return err
	}

	// interfaceFile := fmt.Sprintf("%s/interface.go", dirName)
	// t, err = template.New("interface").Parse(tmp.TmpInterface)
	// if err != nil {
	// 	return err
	// }
	// lf, err = os.Create(interfaceFile)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	lf.Close()
	// }()
	// err = t.Execute(lf, b)
	// if err != nil {
	// 	return err
	// }

	return nil
}
