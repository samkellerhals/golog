package utils

import (
	"fmt"

	"github.com/sdomino/scribble"
)

const dbDir string = ".db"

type JsonDB struct {
	dir    string
	driver *scribble.Driver
}

func (db *JsonDB) init(dir string) *JsonDB {
	driver, err := scribble.New(dir, &scribble.Options{})

	if err != nil {
		fmt.Println(err)
	}

	db.dir = dir
	db.driver = driver
	return db
}

var dbPtr = &JsonDB{}

var InitDB = dbPtr.init(dbDir)
