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

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/nguyen-duc-loc/lch/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	submissionsCmdUsernameFlag = "username"
	colProblemTitle            = "Problem title"
	colSolvedAt                = "Solved at"
)

func submissionsActions(out io.Writer, username string) error {
	lc := leetcode.New()

	submissions, err := lc.Submissions.GetByUsername(username)
	if err != nil {
		return err
	}

	if len(submissions) == 0 {
		_, err = fmt.Fprintln(out, "No recent accepted submissions found")
		return err
	}

	rowHeader := table.Row{
		colProblemTitle,
		colSolvedAt,
	}

	tw := table.NewWriter()
	tw.SetStyle(table.StyleLight)
	tw.SetTitle(utils.BoldText("Recently accepted submissions"))
	tw.AppendHeader(rowHeader)
	tw.Style().Options.SeparateRows = true

	for _, s := range submissions {
		tw.AppendRow(table.Row{
			s.ProblemTitle,
			utils.FormatSince(int64(s.AcTime)),
		})
	}

	_, err = fmt.Fprintln(out, tw.Render())

	return err
}

// submissionsCmd represents the submissions command
var submissionsCmd = &cobra.Command{
	Use:     "submissions",
	Short:   "View recent AC submissions of a user",
	Args:    cobra.NoArgs,
	Aliases: []string{"sm"},
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString(submissionsCmdUsernameFlag)
		if err != nil {
			return err
		}

		if username == "" {
			username = viper.GetString(usernameConfigKey)
		}

		if username == "" {
			return cmd.Usage()
		}

		return submissionsActions(os.Stdout, username)
	},
}

func init() {
	rootCmd.AddCommand(submissionsCmd)

	submissionsCmd.Flags().StringP(submissionsCmdUsernameFlag, "u", "", "username to view profile (default username can be defined in $HOME/.lch.yaml or by using <lch config [flags]>)")
}
