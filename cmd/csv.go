package cmd

import (
	"log"
	"os"
	"time"

	"github.com/Phantas0s/ottosocial/internal"
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
	logpath        string
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
	lp := viper.Get("logpath").(string)
	logger := InitLoggerFile(lp)

	tw, err := internal.NewTwitter(
		viper.Get("key").(string),
		viper.Get("secret").(string),
		viper.Get("token").(string),
		viper.Get("token-secret").(string),
	)
	if err != nil {
		logger.Println(err)
	}

	csv := internal.NewCSV(viper.Get("filepath").(string))
	tweetScheduled, err := csv.Parse()
	if err != nil {
		logger.Println(err)
	}

	if verify {
		errs := tw.VerifyTweetSchedules(tweetScheduled)
		if len(errs) > 0 {
			logger.Println(errs)
			return
		}
	}

	s := gocron.NewScheduler()
	s.Every(1).Second().Do(tw.Sender(tweetScheduled, logger))
	<-s.Start()
}

// TODO makes some of the configuration abstract (via map) in order to reuse it for more commands (?)
func init() {
	RootCmd.AddCommand(csvCmd)
	csvCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Your Twitter Consumer Key (required)")
	csvCmd.PersistentFlags().StringVarP(&consumerSecret, "secret", "s", "", "Your Twitter Consumer Secret (required)")
	csvCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Your Twitter Access Token (required)")
	csvCmd.PersistentFlags().StringVarP(&tokenSecret, "token-secret", "j", "", "Your Twitter Access Token Secret (required)")
	csvCmd.PersistentFlags().StringVarP(&filepath, "filepath", "f", "", "Filepath for your Tweet CSV (required)")
	csvCmd.PersistentFlags().StringVarP(&logpath, "logpath", "l", "", "path for logs")
	csvCmd.PersistentFlags().BoolVarP(&verify, "verify", "v", false, "Verify if the tweets are valid")

	csvCmd.MarkFlagRequired("key")
	csvCmd.MarkFlagRequired("secret")
	csvCmd.MarkFlagRequired("token")
	csvCmd.MarkFlagRequired("token-secret")
	csvCmd.MarkFlagRequired("filepath")

	viper.BindPFlag("key", csvCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("secret", csvCmd.PersistentFlags().Lookup("secret"))
	viper.BindPFlag("token", csvCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("token-secret", csvCmd.PersistentFlags().Lookup("token-secret"))
	viper.BindPFlag("filepath", csvCmd.PersistentFlags().Lookup("filepath"))
	viper.BindPFlag("logpath", csvCmd.PersistentFlags().Lookup("logpath"))
}

func InitLoggerFile(logpath string) *log.Logger {
	if logpath == "" {
		return log.New(os.Stderr, "", 0)
	}

	file, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	l := log.New(file, "", 0)
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05"))

	return l
}
