package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	Env        string
	PostgresDb string
}

func MustLoad() Config {
	//must load ,if not loaded then panic and return
	godotenv.Load(".env")

	// Get port
	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is empty, Its Mendatory")
	}
	// Get Working environment
	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}
	// Get postgres_DB url
	postgresDbUrl := os.Getenv("POSTGRES_DB")
	if postgresDbUrl == "" {
		panic("POSTGRES DB url is Required")
	}

	return Config{
		Port:       port,
		Env:        env,
		PostgresDb: postgresDbUrl,
	}
}
