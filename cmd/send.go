package cmd

import (
	"fmt"

	"github.com/Phantas0s/ottosocial/internal"
	"github.com/spf13/cobra"
)

var (
	message string
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send",
	Run: func(cmd *cobra.Command, args []string) {
		send()
	},
}

// TODO handling error
func send() {
	twitter, err := internal.NewTwitter(
		consumerKey,
		consumerSecret,
		token,
		tokenSecret,
	)
	if err != nil {
		fmt.Println(err)
	}

	twitter.SendTweet(message)
}

func init() {
	RootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&consumerKey, "key", "k", "", "Your Twitter Consumer Key (required)")
	sendCmd.Flags().StringVarP(&consumerSecret, "secret", "s", "", "Your Twitter Consumer Secret (required)")
	sendCmd.Flags().StringVarP(&token, "token", "t", "", "Your Twitter Access Token (required)")
	sendCmd.Flags().StringVarP(&tokenSecret, "tsecret", "j", "", "Your Twitter Access Token Secret (required)")

	sendCmd.MarkFlagRequired("key")
	sendCmd.MarkFlagRequired("secret")
	sendCmd.MarkFlagRequired("token")
	sendCmd.MarkFlagRequired("tsecret")

	sendCmd.Flags().StringVarP(&message, "message", "m", "", "message you want to send")
	sendCmd.MarkFlagRequired("message")
}
