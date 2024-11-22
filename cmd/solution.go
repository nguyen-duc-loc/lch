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
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/nguyen-duc-loc/leetcode-helper/internal/leetcode"
	"github.com/spf13/cobra"
)

func solutionActions(out io.Writer, problemID int32, language string) error {
	lc := leetcode.New()

	solution, err := lc.Solutions.GetByID(problemID, language)
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(out, solution.Content)
	return err
}

// solutionCmd represents the solution command
var solutionCmd = &cobra.Command{
	Use:   "solution <problem_id>",
	Short: "Get the solution to the given problem",
	Long: `Get the solution to the problem <problem_id>.
Or you can get the solution to the today challenge problem by running:
  lch solution today [flags]
	`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}

		language, err := cmd.Flags().GetString("language")
		if err != nil {
			return err
		}

		if args[0] == "today" {
			lc := leetcode.New()
			todayProblem, err := lc.Problems.GetToday()
			if err != nil {
				return err
			}

			return solutionActions(os.Stdout, todayProblem.ID, language)
		}

		problemID, err := strconv.ParseInt(args[0], 10, 32)
		if err != nil || problemID <= 0 || problemID > math.MaxInt32 {
			return fmt.Errorf("problem ID must be an integer in the range of 1 to %d", math.MaxInt32)
		}

		return solutionActions(os.Stdout, int32(problemID), language)
	},
}

func init() {
	rootCmd.AddCommand(solutionCmd)

	extensions := []string{}
	for k := range leetcode.AvailableLanguage {
		extensions = append(extensions, fmt.Sprintf("%q", k))
	}
	solutionCmd.Flags().StringP("language", "l", "", fmt.Sprintf("Choose your favorite language. Available options include: %s", strings.Join(extensions, ", ")))
	solutionCmd.MarkFlagRequired("language")
}
