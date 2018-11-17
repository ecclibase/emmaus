package oksana

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig reads and returns a .env file as a map
func LoadConfig() map[string]string {
	c, err := godotenv.Read()
	if err != nil {
		log.Fatal("ERROR: Could not load .env file")
	}
	return c
}
