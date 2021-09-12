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
		activityDescription := args[2]

		writeToDatabase(activityType, activityDate, activityDescription)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

/* TODO: add interactive commands which will ask for user input determined by configuration file yaml. By default args will
	be filled with default nil value if not specified by user.

E.g. ActivityTypes:
		Hiking:
			MountainName: string
			Latitude: float64
			Longitude: float64
			Location: string
			Route: string
			Elevation: string
			DurationHours: int8
			Description: string
		Climbing:
			RouteName: string
			Lat: string
			Lon: string
			Location: string
			Elevation: string
			ClimbingType: string
			Description: string
*/

// Writes an activity to the database
func writeToDatabase(activityType, activityDate, description string) {
	if err := utils.InitDB.Driver.Write(activityType, activityDate, description); err != nil {
		fmt.Printf("An error occurred during write: %v", err)
	}
	fmt.Println("Succesfully added activity.")
}
