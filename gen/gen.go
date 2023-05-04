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
//allowPkgPath
func Gen(items []GenItem, allowPkgPath []string, outDir string) error {

	b := NewBindData(items, allowPkgPath)
	err := b.Load()
	if err != nil {
		panic(err)
	}
	bs, _ := json.MarshalIndent(b, "", "    ")
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

	for _, obj := range b.AllObj {
		bindFile := fmt.Sprintf("%s/%s_bind.go", dirName, obj.LowerName())
		t, err := template.New("bind").Parse(tmp.TmpBind)
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
	return nil
}
