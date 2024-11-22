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
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/nguyen-duc-loc/leetcode-helper/internal/leetcode"
	"github.com/nguyen-duc-loc/leetcode-helper/utils"
	"github.com/spf13/cobra"
)

const (
	colTitleContestName     = "Contest"
	colTitleRating          = "Rating"
	colTitleFinishTime      = "Finish Time"
	colTitleProblemsSolved  = "Solved"
	colTitleProblemsRanking = "Ranking"
	limitContests           = 15
)

var (
	rowHeader = table.Row{
		colTitleContestName,
		colTitleRating,
		colTitleFinishTime,
		colTitleProblemsSolved,
		colTitleProblemsRanking,
	}
	startRating uint16 = 1500
)

func ratingActions(out io.Writer, username string) error {
	lc := leetcode.New()

	userContestInfo, err := lc.UserContestInfo.GetByUsername(username)
	if err != nil {
		return err
	}

	if userContestInfo.AttendedContestsCount == 0 {
		_, err = fmt.Fprintln(out, "This user hasn't attended any contest before")
		return err
	}

	formattedOut := ""
	formattedOut += outputInfo("Contest attended", fmt.Sprintf("%d", userContestInfo.AttendedContestsCount))
	formattedOut += outputInfo("Rating", leetcode.FormatContestRating(userContestInfo.Rating))
	formattedOut += outputInfo("Global ranking", leetcode.FormatGlobalRanking(userContestInfo.GlobalRanking))
	formattedOut += outputInfo("Contest level", leetcode.FormatContestLevel(userContestInfo.ContestLevel))

	attendedContests := userContestInfo.AttendedContests
	attendedContests = append(attendedContests, leetcode.Contest{
		Rating: startRating,
	})
	if len(attendedContests) >= limitContests {
		attendedContests = attendedContests[:limitContests+1]
	}

	tw := table.NewWriter()
	tw.SetTitle(utils.BoldText("Recently attended contests"))
	tw.AppendHeader(rowHeader)
	tw.SetColumnConfigs([]table.ColumnConfig{
		{Number: 3, Align: text.AlignCenter},
		{Number: 4, Align: text.AlignCenter},
	})

	for i := 0; i < len(attendedContests)-1; i++ {
		contest := attendedContests[i]
		ratingDiff := int16(contest.Rating) - int16(attendedContests[i+1].Rating)
		ratingDiffInText := fmt.Sprintf("%d", ratingDiff)
		if ratingDiff > 0 {
			ratingDiffInText = "+" + ratingDiffInText
		}
		if ratingDiff >= 0 {
			ratingDiffInText = utils.GreenText(ratingDiffInText)
		} else {
			ratingDiffInText = utils.RedText(ratingDiffInText)
		}

		tw.AppendRow(table.Row{
			contest.Metadata.Title,
			fmt.Sprintf("%s (%s)", leetcode.FormatContestRating(contest.Rating), ratingDiffInText),
			utils.FormatTime(int64(contest.FinishTimeInSeconds)),
			fmt.Sprintf("%d/%d", contest.ProblemsSolved, contest.Metadata.TotalProblems),
			contest.Ranking,
		})
	}

	formattedOut += tw.Render()
	_, err = fmt.Fprintln(out, formattedOut)
	return err
}

// ratingCmd represents the rating command
var ratingCmd = &cobra.Command{
	Use:   "rating",
	Short: "Get information about a user's contest rating and his (her) recently attended contests",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString("username")
		if err != nil {
			return err
		}

		return ratingActions(os.Stdout, username)
	},
}

func init() {
	rootCmd.AddCommand(ratingCmd)

	ratingCmd.Flags().StringP("username", "u", "", "username to view rating contest and recently attended contests")
	ratingCmd.MarkFlagRequired("username")
}
