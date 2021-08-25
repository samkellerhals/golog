package utils

import (
	"fmt"

	"github.com/sdomino/scribble"
)

const dbDir string = ".db"

type JsonDB struct {
	Directory string
	Driver    *scribble.Driver
}

func (db *JsonDB) init(dir string) *JsonDB {
	driver, err := scribble.New(dir, &scribble.Options{})

	if err != nil {
		fmt.Println(err)
	}

	db.Directory = dir
	db.Driver = driver
	return db
}

var dbPtr = &JsonDB{}

var InitDB = dbPtr.init(dbDir)
