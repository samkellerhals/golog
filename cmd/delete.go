package cmd

import (
	"fmt"

	"github.com/samkellerhals/golog/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete activity_type activity_date",
	Short: "Deletes an activity from the database.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		activityType := args[0]
		activityDate := args[1]
		deleteFromDatabase(activityType, activityDate)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

// Deletes an activity from the database
func deleteFromDatabase(activityType, activityDate string) {
	if err := utils.InitDB.Driver.Delete(activityType, activityDate); err != nil {
		fmt.Println("An error occured", err)
	}
	fmt.Println("Succesfully deleted activity.")
}
