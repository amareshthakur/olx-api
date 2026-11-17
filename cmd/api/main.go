package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/amareshthakur/olx-api/internal/config"
	"github.com/amareshthakur/olx-api/internal/db"
	"github.com/amareshthakur/olx-api/internal/handlers"
	// _ "net/http/pprof"
)

func main() {
	cfg := config.MustLoad()
	// Connect database
	db, err := db.Connect(cfg.PostgresDb)
	if err != nil {
		log.Fatalf("Main.db.connect: %v", err)
	}
	fmt.Println("Connected to dataBase :-> ", db)
	fmt.Println("Olx server starting", cfg.Port)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)

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
