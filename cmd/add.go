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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new activity to the activity database.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			panic("add requires at least one argument.")
		} else {
			addToDb(args)
		}
	},
}

var (
	ActivityType string
	ActivityDate string
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&ActivityType, "type", "t", "", "set acticity type")
	addCmd.Flags().StringVarP(&ActivityDate, "date", "d", "", "set activity date")
	addCmd.MarkFlagRequired("date")
	addCmd.MarkFlagRequired("type")
}

// Writes a string to the database
func addToDb(args []string) {
	if err := utils.InitDB.Driver.Write(ActivityType, ActivityDate, args); err != nil {
		fmt.Printf("An error occurred during write: %v", err)
	}
}
