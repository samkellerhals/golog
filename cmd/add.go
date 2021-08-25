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
	Short: "Add a new activity.",
	//Long: ``,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().StringVarP(&ActivityType, "type", "t", "", "set acticity type (*)")
	addCmd.MarkFlagRequired("type")

	addCmd.Flags().StringVarP(&ActivityDate, "date", "d", "", "set activity date (*)")
	addCmd.MarkFlagRequired("date")
}

func addToDb(args []string) {

	//fmt.Println(args, Collection, Resource)

	if err := utils.InitDB.Driver.Write(ActivityType, ActivityDate, args); err != nil {
		fmt.Printf("An error occurred during write: %v", err)
	}
}
