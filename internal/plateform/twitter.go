// Twitter wrap the go-twitter package.

package plateform

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
	client *twitter.Client
}

func NewTwitterClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*Twitter, error) {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	_, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	return &Twitter{
		client: client,
	}, nil
}

func (t Twitter) SendTweet(text string) (*twitter.Tweet, error) {
	sentTweet, _, err := t.client.Statuses.Update(text, nil)
	if err != nil {
		return nil, err
	}

	return sentTweet, nil
}

func (t Twitter) ReplyTweet(text string, tweet twitter.Tweet) (*twitter.Tweet, error) {
	sentTweet, _, err := t.client.Statuses.Update(text, &twitter.StatusUpdateParams{InReplyToStatusID: tweet.ID})
	if err != nil {
		return nil, err
	}

	return sentTweet, nil
}

func (t Twitter) SendThread(tweets []string) ([]twitter.Tweet, error) {
	tweetSents := []twitter.Tweet{}
	for k, v := range tweets {
		if len(tweetSents) == 0 {
			tweetSent, err := t.SendTweet(v)
			if err != nil {
				return tweetSents, err
			}
			tweetSents = append(tweetSents, *tweetSent)
		} else {
			tweetSent, err := t.ReplyTweet(v, tweetSents[k-1])
			if err != nil {
				return tweetSents, err
			}
			tweetSents = append(tweetSents, *tweetSent)
		}
	}

	return tweetSents, nil
}
