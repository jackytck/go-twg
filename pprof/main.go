package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// pprof called http.HandleFunc in init()
	// and http.HandleFunc used DefaultServeMux
	// and nil below used DefaultServeMux
	http.ListenAndServe(":3000", nil)
}
