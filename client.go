package anki

import (
	"database/sql"
	"errors"
	"fmt"

	gorp "gopkg.in/gorp.v1"
)

type SQLSelector interface {
	Select(i interface{}, query string, args ...interface{}) ([]interface{}, error)
}

type Client struct {
	DBPath   string
	DBHandle SQLSelector
}

func NewClient(path string) (*Client, error) {
	c := Client{}

	dbObj, err := sql.Open("sqlite3", path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, errors.New("could not connect to Anki DB")
	}

	dbmap := &gorp.DbMap{Db: dbObj, Dialect: gorp.SqliteDialect{}}

	c.DBHandle = dbmap

	return &c, nil
}
