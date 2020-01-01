package internal

import (
	"reflect"
	"testing"
	"time"
)

func Test_ParseCSV(t *testing.T) {
	testCases := []struct {
		name     string
		expected []TweetSchedule
		file     CSV
		wantErr  bool
	}{
		{
			name: "happy case",
			expected: []TweetSchedule{
				{
					Date:      time.Date(2019, 1, 1, 11, 49, 0, 0, time.UTC),
					TweetText: "Hello!! How are you?",
				},
				{
					Date:      time.Date(2019, 2, 2, 11, 40, 0, 0, time.UTC),
					TweetText: "Hello Again!",
				},
			},
			file: CSV{
				filepath: "testdata/test.csv",
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := tc.file.Parse()
			if (err != nil) != tc.wantErr {
				t.Errorf("Error '%v' even if wantErr is %t", err, tc.wantErr)
				return
			}

			if tc.wantErr == false && !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Expected %v, actual %v", tc.expected, actual)
			}
		})
	}
}
