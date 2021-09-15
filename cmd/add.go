package cmd

import (
	"fmt"

	"bufio"
	"os"

	"github.com/samkellerhals/golog/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add activity_type activity_date",
	Short: "Adds a new activity to the activity database.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		activityType := args[0]
		activityDate := args[1]

		activityInformation := getActivityInformation(utils.ActivityFields[activityType])

		writeToDatabase(activityType, activityDate, activityInformation)
	},
}

// Gets activity information from the user via stdin and returns a map with the corresponding data.
func getActivityInformation(fields []string) map[string]interface{} {

	activityInformation := make(map[string]interface{}, len(fields))

	for _, k := range fields {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter %v: \n", k)
		text, _ := reader.ReadString('\n')
		activityInformation[k] = text
	}

	return activityInformation
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// Writes an activity to the database
func writeToDatabase(activityType, activityDate string, description interface{}) {
	if err := utils.InitDB.Driver.Write(activityType, activityDate, description); err != nil {
		fmt.Printf("An error occurred during write: %v", err)
	}
	fmt.Println("Succesfully added activity.")
}
