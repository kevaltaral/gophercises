package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Testhandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test entered test handler, this should not happen")
	}
	return http.HandlerFunc(fn)
}

func TestDevM(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	executeRequest("Get", "/panic", devMw(handler))

}

func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	rr := httptest.NewRecorder()
	rr.Result()
	handler.ServeHTTP(rr, req)
	return rr, err
}

func TestDebugAPI(t *testing.T) {
	// Create server using the a router initialized elsewhere. The router
	// can be a Gorilla mux as in the question, a net/http ServeMux,
	// http.DefaultServeMux or any value that statisfies the net/http
	// Handler interface.
	ts := httptest.NewServer(handler())
	defer ts.Close()

	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	tests := []struct {
		name   string
		r      *http.Request
		status int
	}{
		{name: "1: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/keval/go/src/github.com/kevaltaral/gophercises/recover/main.go&line=24", nil), status: 200},
		{name: "2: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/keval/go/src/github.com/kevaltaral/gophercises/recoer/main.go&line=24", nil), status: 500},
		{name: "2: testing get", r: newreq("GET", ts.URL+"/debug?path=/home/keval/go/src/github.com/kevaltaral/gophercises/recover/main.go&line=et", nil), status: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			defer resp.Body.Close()
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != tt.status {
				t.Error("error in debug api")
			}
		})
	}
}

func TestMakeLinks(t *testing.T) {
	str := `
	goroutine 6 [running]:
runtime/debug.Stack(0xc420049b48, 0x1, 0x1)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xa7
main.devMw.func1.1(0x9ffcc0, 0xc4201a6000)
	/home/keval/go/src/github.com/kevaltaral/gophercises/recover/main.go:74 +0xac
panic(0x82fc80, 0x9f6fd0)
	/usr/local/go/src/runtime/panic.go:502 +0x229
main.funcThatPanics()
	/home/keval/go/src/github.com/kevaltaral/gophercises/recover/main.go:94 +0x39
main.panicDemo(0x9ffcc0, 0xc4201a6000, 0xc420132000)
	/home/keval/go/src/github.com/kevaltaral/gophercises/recover/main.go:85 +0x20
net/http.HandlerFunc.ServeHTTP(0x92be90, 0x9ffcc0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:1947 +0x44
net/http.(*ServeMux).ServeHTTP(0xc42048f950, 0x9ffcc0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:2337 +0x130
main.devMw.func1(0x9ffcc0, 0xc4201a6000, 0xc420132000)
	/home/keval/go/src/github.com/kevaltaral/gophercises/recover/main.go:81 +0x95
net/http.HandlerFunc.ServeHTTP(0xc4204da240, 0x9ffcc0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:1947 +0x44
net/http.serverHandler.ServeHTTP(0xc4204d6a90, 0x9ffcc0, 0xc4201a6000, 0xc420132000)
	/usr/local/go/src/net/http/server.go:2694 +0xbc
net/http.(*conn).serve(0xc420080140, 0xa00040, 0xc42005e100)
	/usr/local/go/src/net/http/server.go:1830 +0x651
created by net/http.(*Server).Serve
	/usr/local/go/src/net/http/server.go:2795 +0x27b`

	makeLinks(str)
}

func TestPanic(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8000/panic", nil)
	if err != nil {
		t.Fatalf("not able to request %v", err)
	}
	rec := httptest.NewRecorder()
	defer func() {
		if err := recover(); err != nil {

		}
	}()
	panicDemo(rec, req)
	res := rec.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("not expected error in panic %v", res.StatusCode)
	}
}
