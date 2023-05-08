package cmd

import (
	"github.com/chintansakhiya/activity/command"
	"github.com/spf13/cobra"
)

func Init() error {

	rootCmd := &cobra.Command{Use: "erp-station"}
	rootCmd.AddCommand(seedData)

	return rootCmd.Execute()
}

var seedData = &cobra.Command{
	Use:   "data",
	Short: "add data to sql",
	Long:  `this command tack data from csv and add to data base`,

	Run: func(cmd *cobra.Command, args []string) {
		command.GetDATA()
	},
}
