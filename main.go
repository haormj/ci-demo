package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	fmt.Println("Listen on 0.0.0.0:8088")
	if err := http.ListenAndServe("0.0.0.0:8088", mux); err != nil {
		fmt.Println(err)
	}
}
