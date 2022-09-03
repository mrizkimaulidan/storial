package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_HOST        string
	DB_PORT        string
	DB_DATABASE    string
	DB_USERNAME    string
	DB_PASSWORD    string
	APP_PORT       string
	JWT_SECRET_KEY string
}

// Get config based on .env file.
// If .env file not found throwing an fatal error.
func (c *Config) GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed opening .env file", err)
	}

	c.DB_HOST = os.Getenv("DB_HOST")
	c.DB_PORT = os.Getenv("DB_PORT")
	c.DB_DATABASE = os.Getenv("DB_DATABASE")
	c.DB_USERNAME = os.Getenv("DB_USERNAME")
	c.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	c.APP_PORT = os.Getenv("APP_PORT")
	c.JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

	return c
}

func New() *Config {
	return &Config{}
}
