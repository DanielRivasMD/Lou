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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	gitCmd := MakeCmd("git", runGit)
	rootCmd.AddCommand(gitCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runGit(cmd *cobra.Command, args []string) {
	const op = "lou.git"

	// inside git repo
	if _, _, err := domovoi.CaptureExecCmd("git", "rev-parse", "--is-inside-work-tree"); err != nil {
		fmt.Println(chalk.Blue.Color("Not a git repository"))
		return
	}

	// git status
	domovoi.PrintCentered("Status")
	if err := domovoi.ExecCmd("git", "status", "--short"); err != nil {
		horus.CheckErr(
			err,
			horus.WithOp(op),
			horus.WithCategory("GIT_ERROR"),
			horus.WithMessage("failed to run git status"),
			horus.WithExitCode(1),
		)
	}

	// git stash list
	domovoi.PrintCentered("Stash List")
	if err := domovoi.ExecCmd("git", "stash", "list"); err != nil {
		horus.CheckErr(
			err,
			horus.WithOp(op),
			horus.WithCategory("GIT_ERROR"),
			horus.WithMessage("failed to list stashes"),
			horus.WithExitCode(1),
		)
	}

	// git log
	domovoi.PrintCentered("Recent Commits")
	if err := domovoi.ExecCmd(
		"git", "log", "--graph", "--topo-order", "--abbrev-commit",
		"--date=relative", "--decorate", "--all", "--boundary",
		"--pretty=format:%Cgreen%ad %Cred%h%Creset -%C(yellow)%d%Creset %s %C(dim white)%cn%Creset",
		"-10",
	); err != nil {
		horus.CheckErr(
			err,
			horus.WithOp(op),
			horus.WithCategory("GIT_ERROR"),
			horus.WithMessage("failed to show git log"),
			horus.WithExitCode(1),
		)
	}

	domovoi.LineBreaks(true)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
