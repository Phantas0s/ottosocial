// Twitter wrap the go-twitter package.

package plateform

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/pkg/errors"
)

const limitTweet = 280

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

func (t Twitter) SendTweet(text string) error {
	_, _, err := t.client.Statuses.Update(text, nil)
	if err != nil {
		return err
	}

	return nil
}

func VerifyTweetLength(message string) (bool, error) {
	if len(message) > limitTweet {
		return false, errors.Errorf("The message contains more than %d characters: \n \"%s\"", limitTweet, message)
	}

	return true, nil
}
