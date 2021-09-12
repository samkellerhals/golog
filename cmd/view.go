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
