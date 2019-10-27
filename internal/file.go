package internal

import (
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/Phantas0s/tweetwee/internal/plateform"
)

const (
	timeLayout = "2006-01-02 03:04"
	date       = iota
	message
)

type File struct {
	filepath string
	twitter  *plateform.Twitter
}

func NewFile(
	filepath,
	consumerKey,
	consumerSecret,
	accessToken,
	accessTokenSecret string,
) (*File, error) {
	t, err := plateform.NewTwitterClient(
		consumerKey,
		consumerSecret,
		accessToken,
		accessTokenSecret,
	)
	if err != nil {
		return nil, err
	}

	return &File{
		filepath: filepath,
		twitter:  t,
	}, nil
}

func (f File) FetchNextTweet() error {
	ts, err := f.ParseCSV()
	if err != nil {
		return err
	}

	for _, v := range ts {
		// TODO not really reliable - lack of precision
		if v.Date.Format(timeLayout) == time.Now().Format(timeLayout) {
			f.twitter.SendTweet(v.TweetText)
		}
	}

	return nil
}

// TODO write test
func (f File) ParseCSV() ([]TweetSchedule, error) {
	// Open the file
	fcsv, err := os.Open(f.filepath)
	if err != nil {
		return nil, err
	}

	tschedules := []TweetSchedule{}
	r := csv.NewReader(fcsv)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		t, err := time.Parse(timeLayout, record[date])
		if err != nil {
			return nil, err
		}

		tschedules = append(tschedules, TweetSchedule{
			Date:      t,
			TweetText: record[message],
		})
	}

	return tschedules, nil
}
