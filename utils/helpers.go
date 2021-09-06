package utils

import (
	"fmt"
	"os"
)

func PrintDirectories(path string) {
	folders, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range folders {
		fmt.Println(f.Name())
	}
}
