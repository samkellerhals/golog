package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var (
	activityDates    []string
	integerDateSlice []int
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

func GetActivityDates(activityType string) []string {
	activityTypePath := filepath.Join(InitDB.Directory, activityType)
	folders, err := os.ReadDir(activityTypePath)

	if err != nil {
		fmt.Println("Error reading directory.")
	}

	for _, f := range folders {
		date := strings.Replace(f.Name(), ".json", "", -1)
		activityDates = append(activityDates, date)
	}

	return activityDates
}

func PrintActivity(date string, activity map[string]interface{}) {
	color.White("%v\n", date)

	for k, v := range activity {
		color.Blue("%v: ", k)
		color.Yellow("%v\n", v)
	}
}
