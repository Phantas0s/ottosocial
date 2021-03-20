package cmd

import (
	"log"
	"os"
	"time"

	"github.com/Phantas0s/ottosocial/internal"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/cobra"
)

var fpath string
var verify bool

func csvCmd(logger *log.Logger) *cobra.Command {
	csvCmd := &cobra.Command{
		Use:   "csv",
		Short: "csv",
		Run: func(cmd *cobra.Command, args []string) {
			csv(logger)
		},
	}

	csvCmd.Flags().StringVarP(&fpath, "filepath", "f", "", "Filepath for your Tweet CSV (required)")
	csvCmd.Flags().BoolVarP(&verify, "verify", "v", false, "Verify if the tweets are valid")

	return csvCmd
}

func csv(logger *log.Logger) {
	tw, err := internal.NewTwitter(consumerKey, consumerSecret, token, tokenSecret)
	if err != nil {
		logger.Println(err)
	}

	csv := internal.NewCSV(fpath)
	tweetScheduled, err := csv.Parse()
	if err != nil {
		logger.Println(err)
	}

	if verify {
		errs := tw.ValidateTweets(tweetScheduled)
		if len(errs) > 0 {
			for _, v := range errs {
				logger.Println(v)
			}
			return
		}
	}

	s := gocron.NewScheduler()
	s.Every(1).Second().Do(tw.Sender(tweetScheduled, logger))
	<-s.Start()
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
	l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " - ")

	return l
}
