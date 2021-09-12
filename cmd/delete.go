package cmd

import (
	"fmt"

	"github.com/samkellerhals/golog/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an activity.",
	//Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		activityType := args[0]
		date := args[1]
		deleteItem(activityType, date)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteItem(activityType, date string) {
	if err := utils.InitDB.Driver.Delete(activityType, date); err != nil {
		fmt.Println("An error occured", err)
	}
	fmt.Println("Succesfully deleted item.")
}
