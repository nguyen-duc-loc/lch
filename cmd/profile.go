/*
Copyright © 2024 Loc Nguyen <nguyenducloc404@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/nguyen-duc-loc/lch/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	profileCmdUsernameFlag = "username"
)

func outputInfo(field, value string) string {
	if len(value) > 0 {
		return fmt.Sprintf("  %s: %s\n", utils.BoldText(field), value)
	}

	return ""
}

func profileActions(out io.Writer, username string) error {
	lc := leetcode.New()

	profile, err := lc.Profiles.GetByUsername(username)
	if err != nil {
		return err
	}

	formattedOut := ""
	formattedOut += outputInfo("Username", profile.Username)
	formattedOut += outputInfo("Name", profile.RealName)
	formattedOut += outputInfo("Bio", profile.Bio)
	formattedOut += outputInfo("School", profile.School)
	formattedOut += outputInfo("Country", profile.Country)
	formattedOut += outputInfo("Github", profile.SocialURL.Github)
	formattedOut += outputInfo("Linkedin", profile.SocialURL.Linkedin)
	formattedOut += outputInfo("Skills", strings.Join(profile.Skills, ", "))
	formattedOut += outputInfo("Global ranking", leetcode.FormatGlobalRanking(profile.GlobalRanking))
	formattedOut += outputInfo("Contest level", leetcode.FormatContestLevel(profile.ContestLevel))

	_, err = fmt.Fprint(out, formattedOut)
	return err
}

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "View user's profile",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString(profileCmdUsernameFlag)
		if err != nil {
			return err
		}

		if username == "" {
			username = viper.GetString(usernameConfigKey)
		}

		if username == "" {
			return cmd.Usage()
		}

		return profileActions(os.Stdout, username)
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	profileCmd.Flags().StringP(profileCmdUsernameFlag, "u", "", "username to view profile (default username can be defined in $HOME/.lch.yaml or by using <lch config [flags]>)")
}
