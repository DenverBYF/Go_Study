package lua

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testHttpServer *httptest.Server

var dynamicResponses = map[string]string{}

func TestMain(m *testing.M) {
	testHttpServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, ok := dynamicResponses[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			_, err := w.Write([]byte("404 Not Found"))
			if err != nil {
				fmt.Println("Write Fail", err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(response))
		if err != nil {
			fmt.Println("Write Fail", err)
		}
	}))
	exitCode := m.Run()

	testHttpServer.Close()

	os.Exit(exitCode)
}

func TestHttpTimer(t *testing.T) {
	t.Run("happy case", func(t *testing.T) {
		dynamicResponses["/test1"] = "success"
		ht := NewHttpTimer(1*time.Second, Req{
			Url:    testHttpServer.URL + "/test1",
			Method: "GET",
		})

		time.Sleep(2 * time.Second)
		assert.True(t, ht.IsRunning())
		ht.Stop()
		time.Sleep(1 * time.Second)
		assert.Greater(t, ht.GetLastUpdateTime(), int64(0))
		assert.False(t, ht.IsRunning())
		assert.Equal(t, ht.GetRequestCount(), 2)
	})

	t.Run("fail count", func(t *testing.T) {
		ht := NewHttpTimer(1*time.Second, Req{
			Url:    testHttpServer.URL + "/404",
			Method: "GET",
		})

		time.Sleep(2 * time.Second)
		assert.True(t, ht.IsRunning())
		ht.Stop()
		time.Sleep(1 * time.Second)
		assert.Greater(t, ht.GetLastUpdateTime(), int64(0))
		assert.False(t, ht.IsRunning())
		assert.Equal(t, ht.GetRequestCount(), 2)
		assert.Equal(t, ht.GetFailCount(), 2)
	})
}
