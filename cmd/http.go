package cmd

import "github.com/spf13/cobra"

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "http",
	Run:   http,
}

func http(cmd *cobra.Command, args []string) {
}

func init() {
	var ck string

	RootCmd.AddCommand(httpCmd)
	httpCmd.Flags().StringVarP(&ck, "consumer", "c", "", "Your Twitter Credential (required)")
	httpCmd.MarkFlagRequired("consumer")
}
