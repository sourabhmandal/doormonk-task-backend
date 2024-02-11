package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfiguration struct {
	HTTP_PORT int
	DB_NAME   string
	GIN_MODE  string
}

var EnvConfig EnvConfiguration

func LoadEnvConfig() {
	godotenv.Load()
	http_port, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 32)
	if err != nil {
		log.Fatalf("\nUnable to parse port value :: %s\nRecieved value :: %s\n", err, os.Getenv("DB_PORT"))
	}

	// --------------------------------------------------------------------------------------
	// --------------------------------------------------------------------------------------
	EnvConfig.HTTP_PORT = int(http_port)
	EnvConfig.DB_NAME = os.Getenv("DB_NAME")
	// Get the value of the GIN_MODE environment variable
	EnvConfig.GIN_MODE = os.Getenv("GIN_MODE")

}
