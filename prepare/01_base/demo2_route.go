package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/*" {
		sayHello2(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func sayHello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my route!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8000", mux)
}
