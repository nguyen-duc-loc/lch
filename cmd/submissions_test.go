package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubmissionsActions(t *testing.T) {
	testCases := []struct {
		name     string
		username string
	}{
		{
			name:     "UserNotExists",
			username: "thisusercannotexistinleetcode",
		},
		{
			name:     "OK",
			username: "nguyenducloc",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer

			err := submissionsActions(&out, tc.username)
			require.NoError(t, err)
			require.NotEmpty(t, out.String())
		})
	}
}
