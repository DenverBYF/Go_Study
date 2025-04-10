package lua

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"example.com/m/v2/httpx"
	lua "github.com/yuin/gopher-lua"
)

type HttpTimer struct {
	// 锁
	mu sync.Mutex
	// 停止通道
	stopChan chan struct{}
	// http client
	hc http.Client
	// 间隔时间
	Interval time.Duration
	// 请求
	Req Req
	// 响应
	Ret *httpx.Ret
	// 最后更新时间
	LastUpdateTime time.Time
	// 错误
	Err error
}

type Req struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    string
}

func NewHttpTimer(interval time.Duration, req Req) *HttpTimer {
	h := &HttpTimer{
		Interval: interval,
		stopChan: make(chan struct{}),
		hc:       http.Client{Timeout: 3 * time.Second},
		Req:      req,
	}
	go h.run()
	return h
}

func (h *HttpTimer) run() {
	ticker := time.NewTicker(h.Interval)
	fmt.Println("start http timer", h.Interval, h.Req.Url)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.mu.Lock()
			switch h.Req.Method {
			case "GET":
				h.Ret, h.Err = httpx.HttpGet(context.Background(), h.hc, h.Req.Url, nil)
				h.LastUpdateTime = time.Now()
				if h.Err != nil {
					fmt.Println("request failed", h.Err)
					continue
				}
				fmt.Println("request success", h.Ret.StatusCode, h.LastUpdateTime.Unix())
			case "POST":
				h.Ret, h.Err = httpx.HttpPost(context.Background(), h.hc, h.Req.Url, h.Req.Body, h.Req.Headers)
				h.LastUpdateTime = time.Now()
				if h.Err != nil {
					fmt.Println("request failed", h.Err)
					continue
				}
				fmt.Println("request success", h.Ret.StatusCode, h.LastUpdateTime.Unix())
			}
			h.mu.Unlock()
		case <-h.stopChan:
			return
		}
	}
}

func (h *HttpTimer) Stop() {
	close(h.stopChan)
}

func (h *HttpTimer) GetLastUpdateTime() int64 {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.LastUpdateTime.Unix()
}

func registerHttpTimer(L *lua.LState) {
	mt := L.NewTypeMetatable("http_timer")
	L.SetGlobal("http_timer", mt)
	methods := map[string]lua.LGFunction{
		"stop": func(L *lua.LState) int {
			h := L.CheckUserData(1).Value.(*HttpTimer)
			h.Stop()
			return 0
		},
		"get_last_update_time": func(L *lua.LState) int {
			h := L.CheckUserData(1).Value.(*HttpTimer)
			L.Push(lua.LNumber(h.GetLastUpdateTime()))
			return 1
		},
	}
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), methods))
}

func createHttpTimer(L *lua.LState) int {
	interval := L.CheckNumber(1)
	timer := NewHttpTimer(time.Duration(interval)*time.Second, Req{
		Url:    "https://www.baidu.com",
		Method: "GET",
	})
	ud := L.NewUserData()
	ud.Value = timer
	L.SetMetatable(ud, L.GetTypeMetatable("http_timer"))
	L.Push(ud)
	return 1
}
