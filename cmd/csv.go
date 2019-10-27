package cmd

import (
	"fmt"

	"github.com/Phantas0s/tweetwee/internal"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
)

var (
	consumerKey    string
	consumerSecret string
	token          string
	tokenSecret    string
	filepath       string
)

var csvCmd = &cobra.Command{
	Use:   "csv",
	Short: "csv",
	Run: func(cmd *cobra.Command, args []string) {
		csv()
	},
}

// TODO handling error
func csv() {
	file, err := internal.NewFile(
		filepath,
		consumerKey,
		consumerSecret,
		token,
		tokenSecret,
	)

	if err != nil {
		fmt.Println(err)
	}

	s := gocron.NewScheduler()
	s.Every(1).Minute().Do(file.FetchNextTweet())
	<-s.Start()
}

func init() {

	RootCmd.AddCommand(csvCmd)
	csvCmd.Flags().StringVarP(&consumerKey, "key", "k", "", "Your Twitter Consumer Key (required)")
	csvCmd.Flags().StringVarP(&consumerSecret, "secret", "s", "", "Your Twitter Consumer Secret (required)")
	csvCmd.Flags().StringVarP(&token, "token", "t", "", "Your Twitter Access Token (required)")
	csvCmd.Flags().StringVarP(&tokenSecret, "tsecret", "j", "", "Your Twitter Access Token Secret (required)")

	csvCmd.MarkFlagRequired("key")
	csvCmd.MarkFlagRequired("secret")
	csvCmd.MarkFlagRequired("token")
	csvCmd.MarkFlagRequired("tsecret")

	csvCmd.Flags().StringVarP(&filepath, "filepath", "f", "", "Filepath for your Tweet CSV (required)")
	csvCmd.MarkFlagRequired("filepath")
}
