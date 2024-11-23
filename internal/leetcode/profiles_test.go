package leetcode_test

import (
	"testing"

	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/stretchr/testify/require"
)

func TestGetProfileByUsername(t *testing.T) {
	lc := leetcode.New()

	testCases := []struct {
		name          string
		username      string
		checkResponse func(t *testing.T, profile *leetcode.Profile, err error)
	}{
		{
			name:     "UserNotExists",
			username: "thisusercannotexistinleetcode",
			checkResponse: func(t *testing.T, profile *leetcode.Profile, err error) {
				require.Error(t, err)
				require.EqualError(t, err, leetcode.ErrUserNotExists.Error())
				require.Empty(t, profile)
			},
		},
		{
			name:     "OK",
			username: "nguyenducloc",
			checkResponse: func(t *testing.T, profile *leetcode.Profile, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, profile)
				require.NotEmpty(t, profile.Username)
				require.NotEmpty(t, profile.RealName)
				require.Positive(t, profile.GlobalRanking)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profile, err := lc.Profiles.GetByUsername(tc.username)
			tc.checkResponse(t, profile, err)
		})
	}
}
