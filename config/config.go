package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// database to store the logs
type Config struct {
	// database
	DB_TYPE     string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSLMODE  string

	// notifier
	NOTIFIER_EMAIL string
}

// setup the database connection
func (c *Config) LoadConfig(env string) error {

	var envfile string

	switch env {
	case "local":
		envfile = ".env.local"
	case "development":
		envfile = ".env.development"
	case "production":
		envfile = ".env.production"
	default:
		envfile = ".env"
	}

	// load env
	err := godotenv.Load(envfile)
	if err != nil {
		return errors.New("Error loading " + envfile + ": " + err.Error())
	}

	// set config values

	// Load configuration variables
	c.DB_TYPE = os.Getenv("DB_TYPE")
	if c.DB_TYPE == "" {
		return errors.New("DB_TYPE not set in the env")
	}

	c.DB_HOST = os.Getenv("DB_HOST")
	if c.DB_HOST == "" {
		return errors.New("DB_HOST not set in the env")
	}

	DB_port := os.Getenv("DB_PORT")
	if DB_port == "" {
		log.Printf("DB_PORT not set, setting default port 5432")
		DB_port = "5432"
	}

	// converting DB_port (string) to uint - total ports 0 -> 65535 -> so 16 bit uint with base 10 (0-9)
	dbport, err := strconv.ParseUint(DB_port, 10, 16)
	if err != nil {
		return errors.New("error parsing the DB_PORT to Integer")
	}
	c.DB_PORT = int(dbport)

	c.DB_USER = os.Getenv("DB_USER")
	if c.DB_USER == "" {
		return errors.New("DB_URL not set in the env")
	}

	c.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	if c.DB_PASSWORD == "" {
		return errors.New("DB_PASSWORD not set in the env")
	}

	c.DB_NAME = os.Getenv("DB_NAME")
	if c.DB_NAME == "" {
		return errors.New("DB_NAME not set in the env")
	}

	c.DB_SSLMODE = os.Getenv("DB_SSLMODE")
	if c.DB_SSLMODE == "" {
		return errors.New("DB_SSLMODE not set in the env")
	}

	c.NOTIFIER_EMAIL = os.Getenv("NOTIFIER_EMAIL")
	if c.NOTIFIER_EMAIL == "" {
		return errors.New("NOTIFIER_EMAIL not set in the env")
	}

	return nil
}
