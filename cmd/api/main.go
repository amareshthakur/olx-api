package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("all ok"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
