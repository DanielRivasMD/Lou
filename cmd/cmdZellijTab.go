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
	"fmt"
	"strings"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	tabLayout    = "tab"
	tabTarget    string
	validTabLayouts = map[string]string{
		"tab":     "Default tab layout",
		"explore": "Explore layout",
		"repl":    "REPL layout",
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var tabCmd = &cobra.Command{
	Use:   "tab [path]",
	Short: "Launch a new Zellij tab",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Italic.TextStyle(chalk.Blue.Color("lilith")) + ` tab launches a new Zellij session in the specified directory using one of the available layouts.

Layouts:
  tab      - ` + validTabLayouts["tab"] + `
  explore  - ` + validTabLayouts["explore"] + `
  repl     - ` + validTabLayouts["repl"] + `

Specify --layout to choose a layout (defaults to "tab").`,
	Example: chalk.Cyan.Color("lou") + " tab ~/src/helix --layout explore",

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Args: cobra.MaximumNArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		const op = "tab.cmd"

		// 1) Validate positional arguments (at most one path)
		switch len(args) {
		case 0:
			tabTarget = ""
		case 1:
			tabTarget = args[0]
		default:
			horus.CheckErr(
				fmt.Errorf("too many arguments: %d", len(args)),
				horus.WithOp(op),
				horus.WithMessage("tab takes at most one directory argument"),
			)
		}

		createTab(tabLayout, tabTarget)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(tabCmd)

	tabCmd.Flags().StringVarP(&tabLayout, "layout", "l", tabLayout, "Layout to use [tab, explore, repl]")

	tabCmd.RegisterFlagCompletionFunc("layout",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var out []string
			for name := range validTabLayouts {
				if toComplete == "" || strings.HasPrefix(name, toComplete) {
					out = append(out, name)
				}
			}
			return out, cobra.ShellCompDirectiveNoFileComp
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
