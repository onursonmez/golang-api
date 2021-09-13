package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/qiniu/qmgo"
)

type Connection interface {
	DB() *qmgo.Database
}

type conn struct {
	client *qmgo.Client
}

func NewConnection() Connection {
	var c conn
	var err error
	url := getURL()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: url})
	c.client = client
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) DB() *qmgo.Database {
	return c.client.Database(os.Getenv("DATABASE_NAME"))
}

func getURL() string {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Println("error on load db port from env:", err.Error())
		port = 27017
	}

	//return fmt.Sprintf("mongodb://%s:%d/%s",
	//	os.Getenv("DATABASE_HOST"),
	//	port,
	//	os.Getenv("DATABASE_NAME"))
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		port,
		os.Getenv("DATABASE_NAME"))
}
