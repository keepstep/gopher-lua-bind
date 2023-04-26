package main

import (
	"flag"
	"log"
	"os"

	"github.com/keepstep/gopher-lua-bind/bind"
	lua "github.com/yuin/gopher-lua"
)

var (
	exec = flag.String("execute", "", "execute lua script")
)

func main() {
	flag.Parse()
	state := lua.NewState()
	defer state.Close()
	bind.Preload(state)
	if *exec != `` {
		if err := state.DoFile(*exec); err != nil {
			log.Printf("[ERROR] Error executing file: %v", err)
		} else {
			log.Printf("top: %v", state.GetTop())
			a, b := state.Get(1), state.Get(2)
			log.Printf("result: %v:%v  %v:%v \n", a.Type().String(), a.String(), b.Type().String(), b.String())
			log.Printf("top-1: %v", state.GetTop())
			state.Pop(2)
			log.Printf("top-2: %v", state.GetTop())
		}
	} else {
		log.Println("Target file was not given!")
		os.Exit(0)
	}
}
