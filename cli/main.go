package cli

import (
	"github.com/spf13/cobra"
)

func Init() error {

	rootCmd := &cobra.Command{Use: "isight"}
	rootCmd.AddCommand(syncData)

	return rootCmd.Execute()
}
