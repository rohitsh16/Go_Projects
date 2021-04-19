package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v : %v\n", name, h)
		}
		// fmt.Fprint(w, "%v: %v\n", name, h)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}