package cmd

import (
	"fmt"

	"github.com/Phantas0s/tweetwee/internal"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	consumerKey    string
	consumerSecret string
	token          string
	tokenSecret    string
	filepath       string
	verify         bool
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
	tw, err := internal.NewTwitter(
		viper.Get("key").(string),
		viper.Get("secret").(string),
		viper.Get("token").(string),
		viper.Get("token-secret").(string),
	)
	if err != nil {
		fmt.Println(err)
	}

	csv := internal.NewCSV(filepath)
	tweetScheduled, err := csv.Parse()
	if err != nil {
		fmt.Println(err)
	}

	if verify {
		errs := tw.VerifyTweetSchedules(tweetScheduled)
		if len(errs) > 0 {
			fmt.Println(errs)
			return
		}
	}

	s := gocron.NewScheduler()
	s.Every(1).Second().Do(tw.Sender(tweetScheduled))
	<-s.Start()
}

func init() {
	RootCmd.AddCommand(csvCmd)
	csvCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Your Twitter Consumer Key (required)")
	csvCmd.PersistentFlags().StringVarP(&consumerSecret, "secret", "s", "", "Your Twitter Consumer Secret (required)")
	csvCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Your Twitter Access Token (required)")
	csvCmd.PersistentFlags().StringVarP(&tokenSecret, "tsecret", "j", "", "Your Twitter Access Token Secret (required)")
	csvCmd.PersistentFlags().StringVarP(&filepath, "filepath", "f", "", "Filepath for your Tweet CSV (required)")
	csvCmd.PersistentFlags().BoolVarP(&verify, "verify", "v", false, "Verify if the tweets are valid")

	csvCmd.MarkFlagRequired("key")
	csvCmd.MarkFlagRequired("secret")
	csvCmd.MarkFlagRequired("token")
	csvCmd.MarkFlagRequired("tsecret")
	csvCmd.MarkFlagRequired("filepath")

	viper.BindPFlag("key", csvCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("secret", csvCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("token", csvCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("token-secret", csvCmd.PersistentFlags().Lookup("tsecret"))
}
