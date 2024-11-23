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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	usernameConfigKey          = "username"
	preferredLanguageConfigKey = "preferred_language"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Version: "0.1",
	Use:     "lch",
	Short:   "A command-line tool designed to streamline your Leetcode experience",
	Long: `Leetcode Helper is a powerful command-line tool that empowers you to efficiently navigate and conquer Leetcode challenges. Key features include:

- User profile fetching: Easily retrieve detailed information about any Leetcode user, including their submission history, ratings, and more.

- Problem content retrieval: Effortlessly fetch the problem statement, constraints, and examples for any Leetcode problem.

- Solution Access: Gain access to a curated collection of solutions for various Leetcode problems, written in popular programming languages.
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(homeDir)
	viper.SetConfigName(".lch")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			configFile, _ := os.Create(fmt.Sprintf("%s/.lch.yaml", homeDir))
			defer configFile.Close()

			viper.Set(usernameConfigKey, "")
			viper.Set(preferredLanguageConfigKey, "")
			viper.WriteConfig()
		} else {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}
