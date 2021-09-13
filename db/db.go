package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/qiniu/qmgo"
)

type Connection interface {
	Close()
	DB() *qmgo.Database
}

type conn struct {
	session *qmgo.Client
}

func NewConnection() Connection {
	var c conn
	var err error
	url := getURL()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	session, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: url})
	c.session = session
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	fmt.Println("database disconnected")
	ctx := context.Background()
	c.session.Close(ctx)
}

func (c *conn) DB() *qmgo.Database {
	return c.session.Database(os.Getenv("DATABASE_NAME"))
}

func getURL() string {
	//port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	//if err != nil {
	//	log.Println("error on load db port from env:", err.Error())
	//	port = 27017
	//}
	return "mongodb://user:password@localhost:27017/mydb?authMechanism=SCRAM-SHA-256"
	//return fmt.Sprintf("mongodb://%s:%d/%s",
	//	os.Getenv("DATABASE_HOST"),
	//	port,
	//	os.Getenv("DATABASE_NAME"))
	//return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
	//	os.Getenv("DATABASE_USER"),
	//	os.Getenv("DATABASE_PASS"),
	//	os.Getenv("DATABASE_HOST"),
	//	port,
	//	os.Getenv("DATABASE_NAME")) + "?authSource=mydb"
}
