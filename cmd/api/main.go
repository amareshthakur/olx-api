package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amareshthakur/olx-api/internal/config"
	// _ "net/http/pprof"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println("Olx server starting", cfg.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		// Order matter, first set header flush it then response body
		w.Header().Set("Content-Type", "application/json") //1
		// w.Write([]byte("all ok"))
		w.WriteHeader(http.StatusOK)       //2
		w.Write([]byte(`{"status":"ok"}`)) //3
	})

	srv := http.Server{
		// Addr:         ":9090",
		// Addr:         ":" + os.Getenv("PORT"),
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	// Max i.e router is a handler which can store multiple handler
	// err := http.ListenAndServe(":9090", mux)
	// log.Print("Server runing on port ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
