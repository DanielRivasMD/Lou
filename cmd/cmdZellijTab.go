/*
Copyright © 2025 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	tabTarget    string
	tabLayout    string
	validLayouts = map[string]string{
		"tab":     "Default tab layout",
		"explore": "Explore layout",
		"repl":    "REPL layout",
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// tabCmd unifies launching tab, explore, and repl layouts.
var tabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Launch a new Zellij tab",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Launch a Zellij tab with one of the following layouts:

  tab      - ` + validLayouts["tab"] + `
  explore  - ` + validLayouts["explore"] + `
  repl     - ` + validLayouts["repl"] + `

Pass --layout to choose one, or omit to use “tab”.
`,
	Example: chalk.Cyan.Color("lou") + " tab --layout explore --target ~/code",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if tabLayout == "" {
			tabLayout = "tab"
		}
		createTab(tabLayout)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(tabCmd)

	tabCmd.Flags().StringVarP(&tabTarget, "target", "t", "", "Change to this directory before launching")
	tabCmd.Flags().StringVarP(&tabLayout, "layout", "l", "", "Layout to use [tab, explore, repl]")

	tabCmd.RegisterFlagCompletionFunc("layout", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var out []string
		for name := range validLayouts {
			if toComplete == "" || name == toComplete || startsWith(name, toComplete) {
				out = append(out, name)
			}
		}
		return out, cobra.ShellCompDirectiveNoFileComp
	})
}

func startsWith(s string, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

////////////////////////////////////////////////////////////////////////////////////////////////////
