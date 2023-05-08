package cli

import (
	"github.com/chintansakhiya/activity/command"
	"github.com/spf13/cobra"
)

var syncData = &cobra.Command{
	Use:   "sync",
	Short: "add data to sql",
	Long:  `this command tack data from csv and add to data base`,

	Run: func(cmd *cobra.Command, args []string) {
		command.GetDATA()
	},
}
