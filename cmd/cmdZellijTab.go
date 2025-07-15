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

// tabTarget is the directory to cd into before creating a tab.
var (
	tabTarget string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// tabCmd is the root-level 'lou tab' command.
var tabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Launch a new Zellij tab with a custom layout",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Lou tab spins up a new Zellij tab using your default layout (` + chalk.Yellow.Color("tab.kdl") + `) 
and names it after the current directory (or "~" if you're in $HOME).`,
	Example: chalk.Cyan.Color("lou") + " tab --target /path/to/dir",
	Run: func(cmd *cobra.Command, args []string) {
		createTab("tab")
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(tabCmd)
	tabCmd.Flags().
		StringVarP(&tabTarget, "target", "t", "", "Change to this directory before launching the tab")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
