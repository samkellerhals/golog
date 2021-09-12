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
