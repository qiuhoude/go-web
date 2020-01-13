package main

import (
	"github.com/panjf2000/ants/v2"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Param  []byte
	Result chan []byte
}

func poolfunc(payload interface{}) {
	request, ok := payload.(*Request)
	if !ok {
		return
	}
	reverseParam := func(s []byte) []byte {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}(request.Param)

	request.Result <- reverseParam
}
func main() {

	//ants.NewPool(100)
	pool, _ := ants.NewPoolWithFunc(1, poolfunc, ants.WithPreAlloc(true))

	defer pool.Release()
	//pool.Tune(10)
	http.HandleFunc("/reverse", func(w http.ResponseWriter, r *http.Request) {
		param, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "request error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		request := &Request{Param: param, Result: make(chan []byte)}

		// Throttle the requests traffic with ants pool. This process is asynchronous and
		// you can receive a result from the channel defined outside.
		if err := pool.Invoke(request); err != nil {
			http.Error(w, "throttle limit error", http.StatusInternalServerError)
		}

		w.Write(<-request.Result)
	})
	_ = http.ListenAndServe(":8080", nil)
}
