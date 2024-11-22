package leetcode_test

import (
	"testing"

	"github.com/nguyen-duc-loc/leetcode-helper/internal/leetcode"
	"github.com/stretchr/testify/require"
)

func TestGetUserContestInfoByUsername(t *testing.T) {
	lc := leetcode.New()

	testCases := []struct {
		name          string
		username      string
		checkResponse func(t *testing.T, userContestInfo *leetcode.UserContestInfo, err error)
	}{
		{
			name:     "UserNotExists",
			username: "thisusercannotexistinleetcode",
			checkResponse: func(t *testing.T, userContestInfo *leetcode.UserContestInfo, err error) {
				require.Error(t, err)
				require.EqualError(t, err, leetcode.ErrUserNotExists.Error())
				require.Empty(t, userContestInfo)
			},
		},
		{
			name:     "OK",
			username: "nguyenducloc",
			checkResponse: func(t *testing.T, userContestInfo *leetcode.UserContestInfo, err error) {
				require.NoError(t, err)
				require.NotEmpty(t, userContestInfo)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			profile, err := lc.UserContestInfo.GetByUsername(tc.username)
			tc.checkResponse(t, profile, err)
		})
	}
}
