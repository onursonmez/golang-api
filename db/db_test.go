package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestGetURL(t *testing.T) {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Println("error on load db port from env:", err.Error())
		port = 27017
	}
	//return fmt.Sprintf("mongodb://%s:%d/%s",
	//	os.Getenv("DATABASE_HOST"),
	//	port,
	//	os.Getenv("DATABASE_NAME"))
	fmt.Printf("mongodb://%s:%s@%s:%d/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		port,
		os.Getenv("DATABASE_NAME"))
}
