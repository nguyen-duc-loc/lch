package leetcode_test

import (
	"testing"

	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/stretchr/testify/require"
)

func TestGetStatsByUsername(t *testing.T) {
	lc := leetcode.New()

	testCases := []struct {
		name          string
		username      string
		checkResponse func(t *testing.T, stats *leetcode.Stats, err error)
	}{
		{
			name:     "UserNotExists",
			username: "thisusercannotexistinleetcode",
			checkResponse: func(t *testing.T, stats *leetcode.Stats, err error) {
				require.Error(t, err)
				require.EqualError(t, err, leetcode.ErrUserNotExists.Error())
				require.Empty(t, stats)
			},
		},
		{
			name:     "OK",
			username: "nguyenducloc",
			checkResponse: func(t *testing.T, stats *leetcode.Stats, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, stats)

				for _, problem := range stats.Problems {
					require.NotEmpty(t, problem.Difficulty)
					require.Positive(t, problem.Total)
				}

				for _, language := range stats.Languages {
					require.NotEmpty(t, language.Name)
					require.Positive(t, language.ProblemsSolved)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			stats, err := lc.Stats.GetByUsername(tc.username)
			tc.checkResponse(t, stats, err)
		})
	}
}
