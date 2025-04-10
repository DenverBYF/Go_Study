package lua

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func LuaTest() {
	fmt.Println("lua test!")
	L := lua.NewState()
	defer L.Close()
	registerHttpTimer(L)
	L.SetGlobal("create_http_timer", L.NewFunction(createHttpTimer))
	err := L.DoString(`
		local http_timer = create_http_timer(1)

		for i = 1, 10 do
			print("lua get last update time", i, http_timer:get_last_update_time())
			os.execute("sleep 3")
		end

		http_timer:stop()
	`)
	if err != nil {
		fmt.Println("load file failed", err)
		return
	}

}
