package internal

import (
	"time"

	"github.com/Phantas0s/tweetwee/internal/plateform"
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

func (t *Twitter) Sender(ts []TweetSchedule) func() error {
	return func() error {
		for _, v := range ts {
			// TODO not really reliable - lack of precision
			if v.Date.Format(timeLayout) == time.Now().Format(timeLayout) {
				err := t.twitter.SendTweet(v.TweetText)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

func (*Twitter) VerifyTweetSchedules(ts []TweetSchedule) []error {
	errors := []error{}
	for _, v := range ts {
		_, err := plateform.VerifyTweet(v.TweetText)
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
