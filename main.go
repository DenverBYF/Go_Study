package main

import (
	lua "github.com/yuin/gopher-lua"
	"strconv"
)

func GoDouble(L *lua.LState) int {
	lv := L.ToInt(1) /* get argument */
	lv2 := lv * 2
	lvs := strconv.Itoa(lv) + " * 2 = " + strconv.Itoa(lv2)
	L.Push(lua.LString(lvs)) /* push result */
	return 1                 /* number of results */
}

func main() {

	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("double", L.NewFunction(GoDouble))
	if err := L.DoFile("test.lua"); err != nil {
		panic(err)
	}
}
