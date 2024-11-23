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
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/nguyen-duc-loc/lch/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	statsCmdUsernameFlag               = "username"
	colProblemDifficulty               = "Difficulty"
	colSolvedProblemsCount             = "Solved"
	colTotalProblemsCount              = "Total"
	colBeat                            = "Beat"
	colLanguageName                    = "Name"
	colSolvedProblemsWithLanguageCount = "Count"
)

func statsActions(out io.Writer, username string) error {
	lc := leetcode.New()

	stats, err := lc.Stats.GetByUsername(username)
	if err != nil {
		return err
	}

	formattedOutput := ""

	problemsSolvedTw := table.NewWriter()
	problemsSolvedTw.SetStyle(table.StyleLight)
	problemsSolvedTw.SetTitle(utils.BoldText("Problems solved"))
	problemsSolvedTw.AppendHeader(table.Row{
		colProblemDifficulty,
		colSolvedProblemsCount,
		colTotalProblemsCount,
		colBeat,
	})

	var solved uint64 = 0
	var total uint64 = 0
	var coloredDifficulty = map[string]string{
		"Easy":   utils.GreenText("Easy"),
		"Medium": utils.YellowText("Medium"),
		"Hard":   utils.RedText("Hard"),
	}

	for _, problem := range stats.Problems {
		solved += problem.ProblemsSolved
		total += problem.Total

		problemsSolvedTw.AppendRow(table.Row{
			coloredDifficulty[problem.Difficulty],
			problem.ProblemsSolved,
			problem.Total,
			fmt.Sprintf("%.2f%%", problem.Beat),
		})
	}

	problemsSolvedTw.AppendFooter(table.Row{
		"Total",
		solved,
		total,
	})

	languagesTw := table.NewWriter()
	languagesTw.SetStyle(table.StyleLight)
	languagesTw.SetTitle(utils.BoldText("Languages used"))
	languagesTw.AppendHeader(table.Row{
		colLanguageName,
		colSolvedProblemsWithLanguageCount,
	})

	languages := stats.Languages
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].ProblemsSolved > languages[j].ProblemsSolved
	})

	for _, language := range languages {
		languagesTw.AppendRow(table.Row{
			language.Name,
			language.ProblemsSolved,
		})
	}

	formattedOutput += problemsSolvedTw.Render()
	formattedOutput += "\n"
	formattedOutput += languagesTw.Render()
	formattedOutput += "\n"

	_, err = fmt.Fprint(out, formattedOutput)

	return err
}

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Get statistics about a user's solving problem process",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := cmd.Flags().GetString(statsCmdUsernameFlag)
		if err != nil {
			return err
		}

		if username == "" {
			username = viper.GetString(usernameConfigKey)
		}

		if username == "" {
			return cmd.Usage()
		}

		return statsActions(os.Stdout, username)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	statsCmd.Flags().StringP(statsCmdUsernameFlag, "u", "", "username to view stats (default username can be defined in $HOME/.lch.yaml or by using <lch config [flags]>)")
}
