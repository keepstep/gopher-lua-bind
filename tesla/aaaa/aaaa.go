package aaaa

import (
	"errors"
	"fmt"
	"net/http"
)

type AAA struct {
	Name      string
	Age       int
	Weight    float32
	Gender    bool
	B         *BBB
	C         CCC
	M         map[string]*CCC
	S         []*BBB
	A         [3]*int
	client    http.Client
	MM        map[string]*http.Client
	lowerCase string
}

func (*AAA) Do(a int, p *BBB, c *CCC) (name string, aaa *AAA, bbb *BBB, err error) {
	return "name", &AAA{}, nil, errors.New("error do")
}
func (*AAA) Call(a int, p *BBB, h *http.Client) (name string, aaa AAA, ccc *CCC, err error) {
	return "", AAA{}, nil, nil
}
func (*AAA) DoAdd(a int, b float32) (r int, err error) {
	return a + int(b), errors.New("error add")
}
func (*AAA) DoMap(a int, b float32, m map[string]int) (r int, f float64, mm map[string]string, err error) {
	mm = map[string]string{}
	for k, v := range m {
		mm[k] = fmt.Sprintf("-%d-", v)
	}
	return a * 2, float64(b * 2), mm, errors.New("error add")
}

func (*AAA) DoSlice(s []int) (ss []string, err error) {
	ss = []string{}
	for _, v := range s {
		ss = append(ss, fmt.Sprintf("-%d-", v))
	}
	return ss, errors.New("error add")
}
func (*AAA) DoSliceSlice(s []int) (ss [][]string, dd [][]float32, mm []map[string]int, err error) {
	ss = [][]string{}
	dd = [][]float32{}
	mm = []map[string]int{}
	return ss, dd, mm, errors.New("error add")
}
func (*AAA) DoMapMap(a int, b float32, m map[string]int) (r int, f float64, mm map[string]map[string]int, ms map[string][]string, err error) {
	mm = map[string]map[string]int{}
	ms = map[string][]string{}
	return a * 2, float64(b * 2), mm, ms, errors.New("error add")
}

func (*AAA) DoZZZ(a int, b any, m map[string]any) any {
	return 2
}

func (*AAA) DoMapAny(a int, b any, m map[string]any) map[string]any {
	return map[string]any{
		"123":  123,
		"abc":  "cba",
		"any1": a,
		"any2": b,
		"any3": m,
	}
}

type BBB struct {
	Tag float64
}

type CCC struct {
	Code int64
}

func AAACmp(name any, dd []any, gg map[int]any) any {
	return nil
}
