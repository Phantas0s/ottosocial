package plateform

import "testing"

func Test_VerifyTweetLength(t *testing.T) {
	testCases := []struct {
		name     string
		expected bool
		message  string
		wantErr  bool
	}{
		{
			name:     "happy case",
			expected: true,
			message:  "hello",
			wantErr:  false,
		},
		{
			name:     "Message too long",
			expected: false,
			message:  "Hello I'm writing quite a lot here maybe a bit too much because it's a tweet after all and I can't write so much stuff so what can I do??? I want to speak more 'cause what I'm saying is quite good, isn't it??? I don't know, maybe Facebook is better? Here's a link: https://thevaluable.dev/abstraction-software-development/",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := VerifyTweetLength(tc.message)
			if (err != nil) != tc.wantErr {
				t.Errorf("Error '%v' even if wantErr is %t", err, tc.wantErr)
				return
			}

			if tc.wantErr == false && actual != tc.expected {
				t.Errorf("Expected %v, actual %v", tc.expected, actual)
			}
		})
	}
}
