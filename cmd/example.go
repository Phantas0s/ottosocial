package cmd

import "github.com/spf13/cobra"

var exampleCmd = &cobra.Command{
	Use:   "example",
	Short: "example",
	Run: func(cmd *cobra.Command, args []string) {
		example()
	},
}

func example() {
}

func init() {
	RootCmd.AddCommand(exampleCmd)
}
