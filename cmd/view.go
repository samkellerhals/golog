/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/samkellerhals/golog/utils"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View activities.",
	Long:  `By default all activities are listed in chronological order, alternatively you can filter by field using the --filter-by option.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Pass at least one argument see golog view -h for more information.")
		} else {
			viewRecords(args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

// Prints all available activity types
func printAllActivityTypes() {
	utils.PrintDirectories(utils.InitDB.Directory)
}

// Prints all logged activities for a specific type
func printActivities(activityType string) {
	records, err := utils.InitDB.Driver.ReadAll(activityType)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range records {
		fmt.Println(string(v))
	}
}

// View activities
func viewRecords(activityType string) {

	if activityType == "all" {
		printAllActivityTypes()
	} else {
		printActivities(activityType)
	}
}
