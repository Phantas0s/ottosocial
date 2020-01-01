package internal

import (
	"encoding/csv"
	"io"
	"os"
	"time"
)

const (
	date = iota
	message
	timeLayout = "2006-01-02 15:04"
)

type CSV struct {
	filepath string
}

func NewCSV(filepath string) *CSV {
	return &CSV{
		filepath: filepath,
	}
}

func (c *CSV) Parse() ([]TweetSchedule, error) {
	// Open the file
	fcsv, err := os.Open(c.filepath)
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
