package leetcode_test

import (
	"testing"

	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/stretchr/testify/require"
)

func TestGetRecentSubmissionsByUsername(t *testing.T) {
	lc := leetcode.New()

	testCases := []struct {
		name          string
		username      string
		checkResponse func(t *testing.T, submissions []*leetcode.Submission, err error)
	}{
		{
			name:     "UserNotExists",
			username: "thisusercannotexistinleetcode",
			checkResponse: func(t *testing.T, submissions []*leetcode.Submission, err error) {
				require.NoError(t, err)
				require.Empty(t, submissions)
			},
		},
		{
			name:     "OK",
			username: "nguyenducloc",
			checkResponse: func(t *testing.T, submissions []*leetcode.Submission, err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			submissions, err := lc.Submissions.GetByUsername(tc.username)
			tc.checkResponse(t, submissions, err)
		})
	}
}
