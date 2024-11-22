package cmd

import (
	"bytes"
	"testing"

	"github.com/nguyen-duc-loc/leetcode-helper/internal/leetcode"
	"github.com/stretchr/testify/require"
)

func TestProfileActions(t *testing.T) {
	testCases := []struct {
		name      string
		username  string
		expectErr error
	}{
		{
			name:      "UserNotExists",
			username:  "thisusercannotexistinleetcode",
			expectErr: leetcode.ErrUserNotExists,
		},
		{
			name:     "OK",
			username: "nguyenducloc",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer

			if err := profileActions(&out, tc.username); err != nil {
				require.Error(t, tc.expectErr)
				require.EqualError(t, err, tc.expectErr.Error())
			} else {
				require.NotEmpty(t, out.String())
			}
		})
	}
}
