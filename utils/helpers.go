package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func PrintActivity(date string, activity interface{}) {
	fmt.Printf("%v\n%v\n\n", date, activity)
}
