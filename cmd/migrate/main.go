package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amareshthakur/olx-api/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		log.Fatal("usage: make migrateUp <up | down>")
	}
	// Load env to conect to DB
	cfg := config.MustLoad()
	m, err := migrate.New(
		// migration file
		"file://migrations",
		// Db connection
		cfg.PostgresDb,
	)
	m.Steps(2)
	switch os.Args[1] {
	case "up":
		log.Printf("migration up")
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	case "down":
		log.Printf("migration down")

		// if err := m.Down(); err != nil {
		if err := m.Steps(-1); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("unknown command: %s", os.Args[1])

	}

}
