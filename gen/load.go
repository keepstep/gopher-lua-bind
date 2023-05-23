package gen

import (
	"fmt"
	"reflect"
	"strings"
)

type BType struct {
	PkgPath string `json:"pkg_path"`
	PkgName string `json:"pkg_name"`
	Name    string `json:"name"`
	IsPtr   bool   `json:"is_ptr"`

	IsBool   bool `json:"is_bool"`
	IsInt    bool `json:"is_int"`
	IsString bool `json:"is_string"`
	IsFloat  bool `json:"is_float"`
	IsNumber bool `json:"is_number"`

	IsStruct bool   `json:"is_struct"`
	IsSlice  bool   `json:"is_slice"`
	IsArray  bool   `json:"is_array"`
	ElemType *BType `json:"elem_type"`

	IsMap       bool   `json:"is_map"`
	ElemKeyType *BType `json:"elem_key_type"`

	IsInterface bool `json:"is_interface"`
	IsError     bool `json:"is_error"`

	IsFunc      bool     `json:"is_func"`
	IsFuncValid bool     `json:"is_func_valid"`
	In          []*BType `json:"in"`
	Out         []*BType `json:"out"`

	RefType reflect.Type `json:"reflect_type"`
	//for fields,methods in out, [1...]
	Index int `json:"index"`
}

type BField struct {
	Name string `json:"name"`
	Type *BType `json:"type"`
}

type BMethod struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	In   []*BType `json:"in"`
	Out  []*BType `json:"out"`
	//全局函数
	IsFunc  bool   `json:"is_func"`
	PkgPath string `json:"pkg_path"`
	PkgName string `json:"pkg_name"`
}

func (t *BType) FuncDefine(pfName string) string {
	pin := []string{}
	pdef := []string{}
	cin := []string{"lp"}
	for i, o := range t.In {
		name := o.Name
		if o.IsStruct {
			name = fmt.Sprintf("* %s.%s", o.PkgName, o.Name)
		}
		pin = append(pin, fmt.Sprintf("ip%d %s", i+1, name))

		s := ""
		if o.IsBool {
			s = fmt.Sprintf("p%d := lua.LBool(ip%d)", i+1, i+1)
		} else if o.IsNumber {
			s = fmt.Sprintf("p%d := lua.LNumber(ip%d)", i+1, i+1)
		} else if o.IsString {
			s = fmt.Sprintf("p%d := lua.LString(ip%d)", i+1, i+1)
		} else if o.IsStruct {
			s = fmt.Sprintf("p%d := Lua_%s_ToUserData(L,ip%d)", i+1, o.Name, i+1)
		} else if o.IsInterface {
			s = fmt.Sprintf("p%d := Lua_%s_Interface_Check(L,ip%d)", i+1, o.Name, i+1)
		}
		pdef = append(pdef, s)
		cin = append(cin, fmt.Sprintf("p%d", i+1))
	}
	rout := []string{}
	rdef := []string{}
	for i, o := range t.Out {
		name := o.Name
		if o.IsStruct {
			name = fmt.Sprintf("* %s.%s", o.PkgName, o.Name)
		}
		rout = append(rout, fmt.Sprintf("r%d %s", i+1, name))
		s := ""
		if o.IsBool {
			s = fmt.Sprintf("r%d = L.CheckBool(top + %d)", i+1, i+1)
		} else if o.IsNumber {
			s = fmt.Sprintf("r%d = %s(L.CheckNumber(top + %d))", i+1, o.Name, i+1)
		} else if o.IsString {
			s = fmt.Sprintf("r%d = L.CheckString(top + %d)", i+1, i+1)
		} else if o.IsStruct {
			s = fmt.Sprintf("r%d = Lua_%s_Check(L,top + %d)", i+1, o.Name, i+1)
		}
		rdef = append(rdef, s)
	}

	str := `func ({PIN}) {ROUT} {
		top := L.GetTop()
		defer func() {
			L.Pop(L.GetTop() - top)
			// fmt.Printf("callback stack change [%d->%d]\n",top,L.GetTop())
		}()
		lp := lua.P{
			Fn:   {PF},
			NRet: {NR},
		}
		{PDEF}
		err := L.CallByParam({CIN})
		if err != nil {
			L.Error(lua.LString(err.Error()), 1)
		}
		{RDEF}
		return
	}`
	pinStr := strings.Join(pin, ",")
	pdefStr := strings.Join(pdef, "\n")
	cinStr := strings.Join(cin, ",")
	routStr := strings.Join(rout, ",")
	if routStr != "" {
		routStr = "(" + routStr + ")"
	}
	rdefStr := strings.Join(rdef, "\n")
	pf := pfName
	nr := fmt.Sprintf("%d", len(t.Out))

	str = strings.ReplaceAll(str, "{PIN}", pinStr)
	str = strings.ReplaceAll(str, "{ROUT}", routStr)
	str = strings.ReplaceAll(str, "{PF}", pf)
	str = strings.ReplaceAll(str, "{NR}", nr)
	str = strings.ReplaceAll(str, "{PDEF}", pdefStr)
	str = strings.ReplaceAll(str, "{CIN}", cinStr)
	str = strings.ReplaceAll(str, "{RDEF}", rdefStr)
	return str
}

func (m *BMethod) InDefine(i int) string {
	if i < 0 || i >= len(m.In) {
		return ""
	}
	t := m.In[i]
	li := i + 2
	if m.IsFunc {
		li = i + 1
	}
	if t.Name == "bool" {
		return fmt.Sprintf("p%d := L.CheckBool(%d)", i+1, li)
	}
	if t.Name == "string" {
		return fmt.Sprintf("p%d := L.CheckString(%d)", i+1, li)
	}
	if t.Name == "int" || t.Name == "uint" {
		return fmt.Sprintf("p%d := L.CheckInt64(%d)", i+1, li)
	}
	if t.Name == "float32" || t.Name == "float64" {
		return fmt.Sprintf("p%d := L.CheckNumber(%d)", i+1, li)
	}
	if t.IsStruct {
		return fmt.Sprintf("p%d := Lua_%s_Check(L,%d)", i+1, t.Name, li)
	}
	if t.IsMap {
		return fmt.Sprintf("p%d := p%dToMap(L,%d)", i+1, i+1, li)
	}
	if t.IsSlice {
		return fmt.Sprintf("p%d := p%dToSlice(L,%d)", i+1, i+1, li)
	}
	if t.IsInterface {
		return fmt.Sprintf("p%d := Lua_%s_Check(L,%d)", i+1, t.Name, li)
	}

	return ""
}
func (m *BMethod) InParam() string {
	c := []string{}
	for i, o := range m.In {
		if strings.HasPrefix(o.Name, "int") || strings.HasPrefix(o.Name, "uint") || strings.HasPrefix(o.Name, "float") {
			c = append(c, fmt.Sprintf("%s(p%d)", o.Name, i+1))
		} else {
			c = append(c, fmt.Sprintf("p%d", i+1))
		}
	}
	return strings.Join(c, ", ")
}
func (m *BMethod) InLen() int {
	return len(m.In)
}
func (m *BMethod) InType() string {
	c := []string{}
	for _, t := range m.In {
		c = append(c, fmt.Sprintf("%s", t.Name))
	}
	return strings.Join(c, ", ")
}

func (m *BMethod) OutRetStr() string {
	c := []string{}
	for i, o := range m.Out {
		if o.IsMap || o.IsSlice {
			c = append(c, fmt.Sprintf("rt%d", i+1))
		} else {
			c = append(c, fmt.Sprintf("r%d", i+1))
		}
	}
	if len(c) > 0 {
		return strings.Join(c, ", ") + " := "
	} else {
		return ""
	}
}

func (m *BMethod) OutCanNil(i int) bool {
	if i < 0 || i >= len(m.Out) {
		return false
	}
	o := m.Out[i]
	if o.IsPtr || o.IsSlice || o.IsArray || o.IsMap || o.IsInterface {
		return true
	} else {
		return false
	}
}

//r1,lua.LNumber(r1)
func (m *BMethod) OutRetArr() [][3]any {
	c := [][3]any{}
	for i, o := range m.Out {
		r := fmt.Sprintf("r%d", i+1)
		s := fmt.Sprintf("r%d", i+1)
		if strings.HasPrefix(o.Name, "int") || strings.HasPrefix(o.Name, "uint") || strings.HasPrefix(o.Name, "float") {
			s = fmt.Sprintf("lua.LNumber(r%d)", i+1)
		} else if o.Name == "string" {
			s = fmt.Sprintf("lua.LString(r%d)", i+1)
		} else if o.Name == "error" && o.IsInterface {
			s = fmt.Sprintf("lua.LString(r%d.Error())", i+1)
		} else if o.IsStruct {
			// ud := fmt.Sprintf("%s_ud", strings.ToLower(o.Name))
			s = fmt.Sprintf("Lua_%s_ToUserData(L,r%d)", o.Name, i+1)
		} else if o.IsInterface {
			// ud := fmt.Sprintf("%s_ud", strings.ToLower(o.Name))
			s = fmt.Sprintf("Lua_%s_ToUserData(L,r%d)", o.Name, i+1)
		} else if o.IsBool {
			s = fmt.Sprintf("lua.LBool(r%d)", i+1)
		} else if o.IsMap {
			r = fmt.Sprintf("r%d", i+1)
		} else if o.IsSlice {
			r = fmt.Sprintf("r%d", i+1)
		}
		c = append(c, [3]any{r, s, o})
	}
	return c
}

func (m *BMethod) OutGet(i int) *BType {
	if i < 0 || i >= len(m.Out) {
		return nil
	}
	return m.Out[i]
}

func (m *BMethod) OutLen() int {
	return len(m.Out)
}

func (m *BMethod) OutType() string {
	c := []string{}
	for _, t := range m.Out {
		c = append(c, fmt.Sprintf("%s", t.Name))
	}
	return strings.Join(c, ", ")
}

type Obj struct {
	BType   `json:"type"`
	Fields  []*BField  `json:"fields"`
	Methods []*BMethod `json:"methods"`
	//global
	Funcs []*BMethod `json:"funcs"`
	//
	Import map[string]int `json:"import"`
}

func (o *Obj) LowerName() string {
	return strings.ToLower(o.Name)
}
func (o *Obj) UdName() string {
	if o.IsInterface {
		return strings.ToLower(o.Name) + "_itf_ud"
	}
	return strings.ToLower(o.Name) + "_ud"
}

// simple field
func (o *Obj) FieldsBind() map[string][2]string {
	m := map[string][2]string{}
	for _, f := range o.Fields {
		t := f.Type
		a := [2]string{"", ""}
		if strings.HasPrefix(t.Name, "int") || strings.HasPrefix(t.Name, "uint") {
			a[0] = fmt.Sprintf("%s(L.CheckInt64(2))", t.Name)
			a[1] = "lua.LNumber"
		} else if strings.HasPrefix(t.Name, "float") {
			a[0] = fmt.Sprintf("%s(L.CheckNumber(2))", t.Name)
			a[1] = "lua.LNumber"
		} else if t.Name == "string" {
			a[0] = fmt.Sprintf("%s(L.CheckString(2))", t.Name)
			a[1] = "lua.LString"
		} else if t.Name == "error" && t.IsInterface {
			a[0] = "errors.New(L.CheckString(2))"
			a[1] = fmt.Sprintf("Lua_%s_ErrorToLv", o.Name)
		} else {
			continue
		}
		m[f.Name] = a
	}
	return m
}

// table field
func (o *Obj) FieldsBindMap() [][3]any {
	m := [][3]any{}
	for _, f := range o.Fields {
		t := f.Type
		if !t.IsMap {
			continue
		}
		if t.ElemKeyType.IsPtr {
			continue
		}
		if !(t.ElemKeyType.IsNumber || t.ElemKeyType.IsString) {
			continue
		}
		if t.ElemType.IsPtr {
			continue
		}
		if !(t.ElemType.IsNumber || t.ElemType.IsString) {
			continue
		}
		a := [3]any{o.Name, f.Name, t}
		m = append(m, a)
	}
	return m
}

func (o *Obj) FieldsBindSlice() [][3]any {
	m := [][3]any{}
	for _, f := range o.Fields {
		t := f.Type
		if !t.IsSlice {
			continue
		}
		if t.ElemType.IsPtr {
			continue
		}
		if !(t.ElemType.IsNumber || t.ElemType.IsString) {
			continue
		}
		a := [3]any{o.Name, f.Name, t}
		m = append(m, a)
	}
	return m
}

func (o *Obj) FieldsBindFunc() [][3]any {
	m := [][3]any{}
	for _, f := range o.Fields {
		t := f.Type
		if !t.IsFunc {
			continue
		}
		if !t.IsFuncValid {
			continue
		}
		a := [3]any{o.Name, f.Name, t}
		m = append(m, a)
	}
	return m
}

//no use，所有xxx_bind 在一个目录
func (o *Obj) GenImportPkg() {
	m := map[string]int{}
	var call func(*BType) = nil
	call = func(b *BType) {
		if b.IsStruct {
			m[b.PkgPath] = 1
		}
		if b.IsMap || b.IsSlice || b.IsArray {
			call(b.ElemType)
		}
	}
	if o.IsStruct {
		m[o.PkgPath] = 1
	}
	// filed ignore
	// for _, t := range o.Fields {
	// 	call(t.Type)
	// }
	for _, m := range o.Methods {
		for _, t := range m.In {
			call(t)
		}
		for _, t := range m.Out {
			call(t)
		}
	}
	if o.Import == nil {
		o.Import = map[string]int{}
	}
	for p := range m {
		o.Import[p] = 1
	}
}

func (o *Obj) AddImport(path string) {
	if o.Import == nil {
		o.Import = map[string]int{}
	}
	o.Import[path] = 1
}

func (o *Obj) AddImportByField(fd *BField) {
	t := fd.Type
	if t.Name == "error" && t.IsInterface {
		o.AddImport("errors")
	}
}

type BindData struct {
	//out
	AllType map[string]*BType `json:"all_type"` //tp.PkgPath()/Name
	AllObj  map[string]*Obj   `json:"all_obj"`  //tp.PkgPath()/Name
	Objs    []*Obj            `json:"objs"`
	//filter
	AllInterface     map[string]*BType `json:"interface"` //tp.PkgPath()/Name
	InterfacePkgPath map[string]int    `json:"interface_pkg_path"`
	AllItf           map[string]*Obj   `json:"all_itf"` //tp.PkgPath()/Name
	//in
	AllowPkgPath map[string]int `json:"allow_pkg_path"`
	Items        []GenItem      `json:"-"`
}

func NewBindData(items []GenItem, allowPkgPath []string) *BindData {
	// if len(allowPkgPath) == 0 {
	// 	allowPkgPath = []string{}
	// 	type em struct{}
	// 	path := reflect.TypeOf(em{}).PkgPath()
	// 	aa := strings.Split(path, "/")
	// 	if len(aa) > 0 {
	// 		allowPkgPath = append(allowPkgPath, aa[0])
	// 	}
	// }
	if allowPkgPath == nil {
		allowPkgPath = []string{}
	}
	// gen AllowPakPath
	m := map[string]int{}
	for _, p := range allowPkgPath {
		m[p] = 1
	}
	for _, item := range items {
		tp := reflect.TypeOf(item.Obj)
		if tp.Kind() == reflect.Pointer {
			tp = tp.Elem()
		}
		if tp.Kind() == reflect.Struct {
			if tp.PkgPath() != "" {
				m[tp.PkgPath()] = 1
			}
		}
	}
	////////////
	return &BindData{
		AllType: map[string]*BType{},
		AllObj:  map[string]*Obj{},
		Objs:    []*Obj{},

		AllInterface:     map[string]*BType{},
		InterfacePkgPath: map[string]int{},
		AllItf:           map[string]*Obj{},

		AllowPkgPath: m,
		Items:        items,
	}
}

func (b *BindData) IsAllowPkgPath(path string) (out bool) {
	if path == "" {
		return true
	}
	for p := range b.AllowPkgPath {
		if strings.HasPrefix(path, p) {
			return true
		}
	}
	return false

}
func (b *BindData) GetBType(name, pkgPath string) (out *BType) {
	key := pkgPath + "/" + name
	if len(b.AllType) == 0 {
		return nil
	}
	out, _ = b.AllType[key]
	return
}
func (b *BindData) CopyBType(t *BType) (out *BType) {
	out = &BType{}
	*out = *t
	return
}
func (b *BindData) AddBType(t *BType) (out *BType) {
	key := t.PkgPath + "/" + t.Name
	out, _ = b.AllType[key]
	if out == nil {
		b.AllType[key] = t
		out = t
	}
	if t.IsInterface {
		outI, _ := b.AllInterface[key]
		if outI == nil {
			b.AllInterface[key] = b.CopyBType(t)
			ikey := t.PkgPath
			if ikey != "" {
				b.InterfacePkgPath[ikey] = 1
			}
		}
	}
	return
}
func (b *BindData) GetObj(name, pkgPath string) (out *Obj) {
	key := pkgPath + "/" + name
	if len(b.AllObj) == 0 {
		return nil
	}
	out, _ = b.AllObj[key]
	return
}
func (b *BindData) AddObj(t *Obj) (out *Obj) {
	key := t.PkgPath + "/" + t.Name
	out, _ = b.AllObj[key]
	if out == nil {
		b.AllObj[key] = t
		out = t
		b.Objs = append(b.Objs, t)
	}
	return
}

func (b *BindData) AddInterface(t *Obj) (out *Obj) {
	key := t.PkgPath + "/" + t.Name
	out, _ = b.AllItf[key]
	if out == nil {
		b.AllItf[key] = t
		out = t
	}
	return
}

// load struct
func (b *BindData) LoadObj(obj any) (*Obj, error) {
	otp := reflect.TypeOf(obj)
	tp := reflect.TypeOf(obj)
	isPtr := false
	if tp.Kind() == reflect.Pointer {
		tp = tp.Elem()
		isPtr = true
	}
	wholeName := fmt.Sprintf("%s/%s", tp.PkgPath(), tp.Name())
	if tp.Kind() != reflect.Struct {
		return nil, fmt.Errorf("load err not struct : %s", wholeName)
	}
	//load type
	btp, ignore := b.LoadType(tp, nil)
	if ignore {
		return nil, fmt.Errorf("load err ignore : %s", wholeName)
	}
	if btp == nil {
		return nil, fmt.Errorf("load err loadType : %s", wholeName)
	}

	fmt.Printf("obj info:%v %v\n", otp.String(), tp.PkgPath())
	fmt.Printf("obj type:%v\n", tp.Name())
	fmt.Printf("obj num_field:%v\n", tp.NumField())
	fmt.Printf("obj num_method:%v\n", otp.NumMethod())

	//get cache
	bobj := b.GetObj(btp.Name, btp.PkgPath)
	if bobj != nil {
		return nil, fmt.Errorf("load err GetObj : %s", wholeName)
	}
	btp.IsPtr = isPtr
	bObj := &Obj{
		BType:   *btp,
		Fields:  []*BField{},
		Methods: []*BMethod{},
	}
	//add to cache
	b.AddObj(bObj)
	//parse filed
	nf := tp.NumField()
	for i := 0; i < nf; i++ {
		ftp := tp.Field(i).Type
		fname := tp.Field(i).Name
		btp, ignore := b.LoadType(ftp, nil)
		if ignore {
			fmt.Printf("ignore field : %s :%s %d %s/%s\n", wholeName, fname, i+1, ftp.Name(), ftp.PkgPath())
			continue
		}
		if btp == nil {
			return nil, fmt.Errorf("load err loadField : %s : %d %s", wholeName, i+1, fname)
		}
		bfd := &BField{
			Name: fname,
			Type: btp,
		}
		if btp.IsFunc {
			err := b.LoadFuncParam(btp)
			if err != nil {
				fmt.Printf("ignore field func err: %s :%s %d %s/%s %s", wholeName, i+1, err.Error())
				continue
			}
		}
		bObj.Fields = append(bObj.Fields, bfd)
		bObj.AddImportByField(bfd)
	}
	if !isPtr {
		return nil, fmt.Errorf("load err isPtr false : %s", wholeName)
	}
	//parse method
	nm := otp.NumMethod()
	for i := 0; i < nm; i++ {
		md := otp.Method(i)
		bmd := &BMethod{
			Name: md.Name,
			Type: md.Type.String(),
			In:   []*BType{},
			Out:  []*BType{},
		}
		ignoreMethod := false
		// skip 0 self
		for i := 1; i < md.Type.NumIn(); i++ {
			t := md.Type.In(i)
			btp, ignore := b.LoadType(t, nil)
			if ignore {
				ignoreMethod = true
				break
			}
			if btp == nil {
				return nil, fmt.Errorf("load err loadMethod In: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsStruct && !btp.IsPtr {
				return nil, fmt.Errorf("load err loadMethod In must be pointer: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsFunc {
				ignoreMethod = true
				break
			}
			btp.Index = i
			bmd.In = append(bmd.In, btp)
		}
		if ignoreMethod {
			fmt.Printf("ignore method : %s :%s %d %s/%s\n", wholeName, md.Type.String(), i+1, tp.Name(), tp.PkgPath())
			continue
		}
		for i := 0; i < md.Type.NumOut(); i++ {
			t := md.Type.Out(i)
			btp, ignore := b.LoadType(t, nil)
			if ignore {
				ignoreMethod = true
				break
			}
			if btp == nil {
				return nil, fmt.Errorf("load err loadMethod Out: %s : %s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsStruct && !btp.IsPtr {
				return nil, fmt.Errorf("load err loadMethod Out must be pointer: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsFunc {
				ignoreMethod = true
				break
			}
			btp.Index = i + 1
			bmd.Out = append(bmd.Out, btp)
		}
		if ignoreMethod {
			fmt.Printf("ignore method : %s :%s %d %s/%s\n", wholeName, md.Type.String(), i+1, tp.Name(), tp.PkgPath())
			continue
		} else {
			bObj.Methods = append(bObj.Methods, bmd)
		}
	}
	// bObj.GenImportPkg()
	return bObj, nil
}

// load interface
func (b *BindData) LoadInterface(btp *BType) (*Obj, error) {
	tp := btp.RefType
	wholeName := fmt.Sprintf("%s/%s", tp.PkgPath(), tp.Name())
	if tp.Kind() != reflect.Interface {
		return nil, fmt.Errorf("load err not interface : %s", wholeName)
	}

	fmt.Printf("itf info:%v %v\n", tp.String(), tp.PkgPath())
	fmt.Printf("itf type:%v\n", tp.Name())
	fmt.Printf("itf num_method:%v\n", tp.NumMethod())

	//get cache
	bobj := b.GetObj(btp.Name, btp.PkgPath)
	if bobj != nil {
		return nil, fmt.Errorf("load itf err GetObj : %s", wholeName)
	}
	bObj := &Obj{
		BType:   *btp,
		Fields:  []*BField{},
		Methods: []*BMethod{},
	}
	//add to cache
	b.AddInterface(bObj)
	//parse method
	nm := tp.NumMethod()
	for i := 0; i < nm; i++ {
		md := tp.Method(i)
		bmd := &BMethod{
			Name: md.Name,
			Type: md.Type.String(),
			In:   []*BType{},
			Out:  []*BType{},
		}
		ignoreMethod := false
		for i := 0; i < md.Type.NumIn(); i++ {
			t := md.Type.In(i)
			btp, ignore := b.LoadType(t, nil)
			if ignore {
				ignoreMethod = true
				break
			}
			if btp == nil {
				return nil, fmt.Errorf("load itf err loadMethod In: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsStruct && !btp.IsPtr {
				return nil, fmt.Errorf("load itf err loadMethod In must be pointer: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsFunc {
				ignoreMethod = true
				break
			}
			btp.Index = i
			bmd.In = append(bmd.In, btp)
		}
		if ignoreMethod {
			fmt.Printf("ignore itf method : %s :%s %d %s/%s\n", wholeName, md.Type.String(), i+1, tp.Name(), tp.PkgPath())
			continue
		}
		for i := 0; i < md.Type.NumOut(); i++ {
			t := md.Type.Out(i)
			btp, ignore := b.LoadType(t, nil)
			if ignore {
				ignoreMethod = true
				break
			}
			if btp == nil {
				return nil, fmt.Errorf("load itf err loadMethod Out: %s : %s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsStruct && !btp.IsPtr {
				return nil, fmt.Errorf("load itf err loadMethod Out must be pointer: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsFunc {
				ignoreMethod = true
				break
			}
			btp.Index = i + 1
			bmd.Out = append(bmd.Out, btp)
		}
		if ignoreMethod {
			fmt.Printf("ignore itf method : %s :%s %d %s/%s\n", wholeName, md.Type.String(), i+1, tp.Name(), tp.PkgPath())
			continue
		} else {
			bObj.Methods = append(bObj.Methods, bmd)
		}
	}
	// bObj.GenImportPkg()
	return bObj, nil
}

// load global func
func (b *BindData) LoadFunc(bObj *Obj, funcs [][2]any) error {

	for _, iii := range funcs {
		name := iii[0].(string)
		fff := iii[1]
		tp := reflect.TypeOf(fff)
		wholeName := fmt.Sprintf("%s/%s", tp.PkgPath(), name)
		if tp.Kind() != reflect.Func {
			return fmt.Errorf("load err not func : %s", wholeName)
		}

		md := &BMethod{
			Name:    name,
			Type:    tp.String(),
			In:      []*BType{},
			Out:     []*BType{},
			IsFunc:  true,
			PkgName: bObj.PkgName,
			PkgPath: bObj.PkgPath,
		}
		ignoreMethod := false
		//
		for i := 0; i < tp.NumIn(); i++ {
			t := tp.In(i)
			btp, ignore := b.LoadType(t, nil)
			if ignore {
				ignoreMethod = true
				break
			}
			if btp == nil {
				return fmt.Errorf("load err LoadFunc In: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsStruct && !btp.IsPtr {
				return fmt.Errorf("load err LoadFunc In must be pointer: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsFunc {
				ignoreMethod = true
				break
			}
			btp.Index = i + 1
			md.In = append(md.In, btp)
		}
		if ignoreMethod {
			fmt.Printf("ignore func : %s :%s %s/%s\n", wholeName, tp.String(), tp.Name(), tp.PkgPath())
			continue
		}
		for i := 0; i < tp.NumOut(); i++ {
			t := tp.Out(i)
			btp, ignore := b.LoadType(t, nil)
			if ignore {
				ignoreMethod = true
				break
			}
			if btp == nil {
				return fmt.Errorf("load err LoadFunc Out: %s : %s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsStruct && !btp.IsPtr {
				return fmt.Errorf("load err LoadFunc Out must be pointer: %s :%s %d %s/%s", wholeName, md.Name, i+1, t.Name(), t.PkgPath())
			}
			if btp.IsFunc {
				ignoreMethod = true
				break
			}
			btp.Index = i + 1
			md.Out = append(md.Out, btp)
		}
		if ignoreMethod {
			fmt.Printf("ignore method : %s :%s %s/%s\n", wholeName, tp.String(), tp.Name(), tp.PkgPath())
			continue
		} else {
			bObj.Funcs = append(bObj.Funcs, md)
		}
	}
	return nil
}

// load func in out
func (b *BindData) LoadFuncParam(ptp *BType) error {
	if !ptp.IsFunc {
		return nil
	}
	ptp.IsFuncValid = false
	tp := ptp.RefType
	ignoreMethod := false
	//
	for i := 0; i < tp.NumIn(); i++ {
		t := tp.In(i)
		btp, ignore := b.LoadType(t, nil)
		if ignore {
			ignoreMethod = true
			break
		}
		if btp == nil {
			return fmt.Errorf("invalid in nil")
		}
		if btp.IsStruct && !btp.IsPtr {
			return fmt.Errorf("invalid in struct not ptr")
		}
		if btp.IsFunc {
			return fmt.Errorf("invalid in is func")
		}
		if !(btp.IsNumber || btp.IsString || btp.IsBool) {
			return fmt.Errorf("invalid in is not number string bool")
		}
		btp.Index = i + 1
		ptp.In = append(ptp.In, btp)
	}
	if ignoreMethod {
		return fmt.Errorf("invalid in ignore")
	}
	for i := 0; i < tp.NumOut(); i++ {
		t := tp.Out(i)
		btp, ignore := b.LoadType(t, nil)
		if ignore {
			ignoreMethod = true
			break
		}
		if btp == nil {
			return fmt.Errorf("invalid out nil")
		}
		if btp.IsStruct && !btp.IsPtr {
			return fmt.Errorf("invalid out struct not ptr")
		}
		if btp.IsFunc {
			return fmt.Errorf("invalid out is func")
		}
		if !(btp.IsNumber || btp.IsString || btp.IsBool) {
			return fmt.Errorf("invalid out is not number string bool")
		}
		btp.Index = i + 1
		ptp.Out = append(ptp.Out, btp)
	}
	if ignoreMethod {
		return fmt.Errorf("invalid out ignore")
	}
	ptp.IsFuncValid = true
	return nil
}

//return value is a copy from AllType
func (b *BindData) LoadType(tp reflect.Type, btpIn *BType) (btp *BType, ignore bool) {
	btp = btpIn
	if btp == nil {
		btp = &BType{}
	}
	kd := tp.Kind()
	if tp.PkgPath() != "" && (kd == reflect.Struct || kd == reflect.Interface) {
		if b.IsAllowPkgPath(tp.PkgPath()) {

		} else {
			return nil, true
		}
	}
	switch kd {
	case reflect.Pointer:
		{
			btp, ignore = b.LoadType(tp.Elem(), btp)
			if btp != nil {
				btp.IsPtr = true
			}
		}
	case reflect.Slice:
		{
			etp, ignoreElem := b.LoadType(tp.Elem(), nil)
			if ignoreElem {
				ignore = true
				break
			}
			if etp == nil {
				btp = nil
			} else {
				// btp.Name = "slice"
				btp.Name = tp.Name()
				btp.PkgPath = tp.PkgPath()
				btp.ElemType = etp
				btp.IsSlice = true
				btp.RefType = tp
				if etp.IsPtr {
					btp.Name = fmt.Sprintf("[]*%s", etp.Name)
				} else {
					btp.Name = fmt.Sprintf("[]%s", etp.Name)
				}
			}
		}
	case reflect.Array:
		{
			etp, ignoreElem := b.LoadType(tp.Elem(), nil)
			if ignoreElem {
				ignore = true
				break
			}
			if etp == nil {
				btp = nil
			} else {
				// btp.Name = "array"
				btp.Name = tp.Name()
				btp.PkgPath = tp.PkgPath()
				btp.ElemType = etp
				btp.IsArray = true
				btp.RefType = tp
				n := tp.Len()
				if etp.IsPtr {
					btp.Name = fmt.Sprintf("[%d]*%s", n, etp.Name)
				} else {
					btp.Name = fmt.Sprintf("[%d]%s", n, etp.Name)
				}
			}
		}
	case reflect.Map:
		{
			ktp, ignoreElem := b.LoadType(tp.Key(), nil)
			if ignoreElem {
				ignore = true
				break
			}
			if ktp == nil {
				btp = nil
				break
			}
			etp, ignoreElem := b.LoadType(tp.Elem(), nil)
			if ignoreElem {
				ignore = true
				break
			}
			if etp == nil {
				btp = nil
				break
			}
			// btp.Name = "map"
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			btp.ElemType = etp
			btp.ElemKeyType = ktp
			btp.IsMap = true
			btp.RefType = tp
			if etp.IsPtr {
				btp.Name = fmt.Sprintf("map[%s]*%s", ktp.Name, etp.Name)
			} else {
				btp.Name = fmt.Sprintf("map[%s]%s", ktp.Name, etp.Name)
			}
		}
	case reflect.Struct:
		{
			//get cache
			out := b.GetBType(tp.Name(), tp.PkgPath())
			if out != nil {
				return b.CopyBType(out), false
			}
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			ss := strings.Split(btp.PkgPath, "/")
			btp.PkgName = ss[len(ss)-1]
			btp.RefType = tp
			btp.IsStruct = true
			btp.IsPtr = false
			//add to cache
			b.AddBType(btp)
			//field type
			nf := tp.NumField()
			for i := 0; i < nf; i++ {
				b.LoadType(tp.Field(i).Type, nil)
			}
			return b.CopyBType(btp), false
		}
	case reflect.Bool:
		{
			//get cache
			out := b.GetBType(tp.Name(), tp.PkgPath())
			if out != nil {
				return b.CopyBType(out), false
			}
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			btp.RefType = tp
			btp.IsBool = true
			//add to cache
			b.AddBType(btp)
		}
	case reflect.String:
		{
			//get cache
			out := b.GetBType(tp.Name(), tp.PkgPath())
			if out != nil {
				return b.CopyBType(out), false
			}
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			btp.RefType = tp
			btp.IsString = true
			//add to cache
			b.AddBType(btp)
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		{
			//get cache
			out := b.GetBType(tp.Name(), tp.PkgPath())
			if out != nil {
				return b.CopyBType(out), false
			}
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			btp.RefType = tp
			btp.IsInt = true
			btp.IsNumber = true
			//add to cache
			b.AddBType(btp)
		}
	case reflect.Float32, reflect.Float64:
		{
			//get cache
			out := b.GetBType(tp.Name(), tp.PkgPath())
			if out != nil {
				return b.CopyBType(out), false
			}
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			btp.RefType = tp
			btp.IsFloat = true
			btp.IsNumber = true
			//add to cache
			b.AddBType(btp)
		}
	case reflect.Interface:
		{
			//error special
			//get cache
			out := b.GetBType(tp.Name(), tp.PkgPath())
			if out != nil {
				return b.CopyBType(out), false
			}
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			ss := strings.Split(btp.PkgPath, "/")
			btp.PkgName = ss[len(ss)-1]
			btp.RefType = tp
			btp.IsInterface = true
			if btp.Name == "error" {
				btp.IsError = true
			}
			//add to cache
			b.AddBType(btp)
		}
	case reflect.Func:
		{
			btp.Name = tp.Name()
			btp.PkgPath = tp.PkgPath()
			btp.RefType = tp
			btp.IsFunc = true
			//when LoadFuncParam
			btp.IsFuncValid = false
			btp.In = []*BType{}
			btp.Out = []*BType{}
		}
	default:
		{
			return nil, true
		}
	}
	if ignore {
		btp = nil
	}
	return btp, ignore
}

// laod all
func (b *BindData) Load() error {
	for _, item := range b.Items {
		obj, err := b.LoadObj(item.Obj)
		if err != nil {
			fmt.Printf("Error Load:%v\n", err)
			return err
		}
		err = b.LoadFunc(obj, item.Funcs)
		if err != nil {
			fmt.Printf("Error LoadFunc:%v\n", err)
			return err
		}
	}
	for _, btp := range b.AllInterface {
		_, err := b.LoadInterface(btp)
		if err != nil {
			fmt.Printf("Error Load Itf:%v\n", err)
			return err
		}
	}
	return nil
}
