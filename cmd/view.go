package cmd

import (
	"fmt"

	"github.com/samkellerhals/golog/utils"
	"github.com/spf13/cobra"
)

var databaseRecord interface{}

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view [activity_type]",
	Short: "View all activities by activity type.",
	Long:  `By default all activity types are listed, to view activities by activity type specify the type.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			printAllActivityTypes()
		} else {
			printByActivityType(args[0])
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
func printByActivityType(activityType string) {

	activityDates := utils.GetActivityDates(activityType)

	for _, date := range activityDates {
		if err := utils.InitDB.Driver.Read(activityType, date, &databaseRecord); err != nil {
			fmt.Println(err)
		}
		utils.PrintActivity(date, databaseRecord)
	}
}
