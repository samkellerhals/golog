package utils

import (
	"fmt"

	"github.com/sdomino/scribble"
)

const dbDir string = ".db"

// Creates a new scribble database
func CreateDb() *scribble.Driver {
	driver, err := scribble.New(dbDir, &scribble.Options{})

	if err != nil {
		fmt.Println(err)
	}
	return driver
}
