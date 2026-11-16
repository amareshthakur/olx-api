package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() Config {
	//must load ,if not loaded then panic and return
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is empty, Its Mendatory")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}
	return Config{
		Port: port,
		Env:  env,
	}
}
