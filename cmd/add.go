package cmd

import (
	"fmt"

	"github.com/samkellerhals/golog/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add activity_type activity_date activity_description",
	Short: "Adds a new activity to the activity database.",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		activityType := args[0]
		activityDate := args[1]
		activityDescription := args[3]

		writeToDatabase(activityType, activityDate, activityDescription)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// Writes a string to the database
func writeToDatabase(activityType, activityDate, description string) {
	if err := utils.InitDB.Driver.Write(activityType, activityDate, description); err != nil {
		fmt.Printf("An error occurred during write: %v", err)
	}
	fmt.Println("Succesfully added activity.")
}
