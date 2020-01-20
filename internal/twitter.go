package internal

import (
	"log"
	"time"

	"github.com/Phantas0s/ottosocial/internal/plateform"
	"github.com/pkg/errors"
)

type Twitter struct {
	twitter *plateform.Twitter
}

type TweetSchedule struct {
	Date      time.Time
	TweetText string
}

func NewTwitter(
	consumerKey,
	consumerSecret,
	accessToken,
	accessTokenSecret string,
) (*Twitter, error) {
	t, err := plateform.NewTwitterClient(
		consumerKey,
		consumerSecret,
		accessToken,
		accessTokenSecret,
	)
	if err != nil {
		return nil, err
	}

	return &Twitter{
		twitter: t,
	}, nil
}

func (t *Twitter) SendTweet(message string) error {
	err := t.twitter.SendTweet(message)
	if err != nil {
		return err
	}

	return nil
}

func (t *Twitter) Sender(ts []TweetSchedule, logger *log.Logger) func() error {
	return func() error {
		defer func() {
			if r := recover(); r != nil {
				logger.Printf("A panic occured: %v \n", r)
			}
		}()
		for _, v := range ts {
			// TODO not really reliable / lack of precision (?)
			now := time.Now().Format(timeLayout + ":05")
			if v.Date.Format(timeLayout+":05") == now {
				err := t.twitter.SendTweet(v.TweetText)
				logger.Printf("The tweet '%s' was sent", v.TweetText)
				if err != nil {
					return errors.Wrapf(err, "Error while sending the message '%s' to Twitter", v.TweetText)
				}
			}
		}

		return nil
	}
}

func (*Twitter) VerifyTweetSchedules(ts []TweetSchedule) []error {
	errors := []error{}
	for _, v := range ts {
		_, err := plateform.VerifyTweetLength(v.TweetText)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
