package main

import (
	"log"
	"net/http"
	// _ "net/http/pprof"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("all ok"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
	// Max i.e router is a handler which can store multiple handler
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
