/*
Copyright © 2024 Daniel Rivas <danielrivasmd@gmail.com>

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

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// gitCmd
var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Get the status of the repository effortlessly",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Blue.Color("lou") + ` simplifies interactions with Git by providing streamlined commands and features
`,
	Example: chalk.White.Color("lou") + " " +
		chalk.Bold.TextStyle(chalk.White.Color("git")),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		const op = "cmd.git"

		// Are we inside a git repo?
		if _, _, err := domovoi.CaptureExecCmd("git", "rev-parse", "--is-inside-work-tree"); err != nil {
			fmt.Println(chalk.Blue.Color("Not a git repository"))
			return
		}

		// Section: git status
		domovoi.PrintCentered("Status")
		if err := domovoi.ExecCmd("git", "status", "--short"); err != nil {
			msg := horus.FormatPanic(op, "failed to run git status")
			fmt.Fprintln(cmd.ErrOrStderr(), msg)
			os.Exit(1)
		}

		// Section: git stash list (capture output)
		domovoi.PrintCentered("Stash List")
		if err := domovoi.ExecCmd("git", "stash", "list"); err != nil {
			msg := horus.FormatPanic(op, "failed to list stashes")
			fmt.Fprintln(cmd.ErrOrStderr(), msg)
			os.Exit(1)
		}

		// Section: git log
		domovoi.PrintCentered("Recent Commits")
		if err := domovoi.ExecCmd(
			"git", "log", "--graph", "--topo-order", "--abbrev-commit",
			"--date=relative", "--decorate", "--all", "--boundary",
			"--pretty=format:%Cgreen%ad %Cred%h%Creset -%C(yellow)%d%Creset %s %C(dim white)%cn%Creset",
			"-10",
		); err != nil {
			msg := horus.FormatPanic(op, "failed to show git log")
			fmt.Fprintln(cmd.ErrOrStderr(), msg)
			os.Exit(1)
		}

		domovoi.LineBreaks(true)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(gitCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
