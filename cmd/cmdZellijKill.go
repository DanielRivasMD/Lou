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

import (
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// zellijKillCmd
var zellijKillCmd = &cobra.Command{
	Use:   "kill",
	Aliases: []string{"k"},
	Short:  "Kill the current Zellij session",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` allow to terminate the current active Zellij session effortlessly by identifying and killing it directly
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("kill") + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// command
		cmdZellijKill := `zellij kill-session "$(zellij list-sessions | grep '(current)' | sed 's/\x1b\[[0-9;]*m//g' | awk '{print $1}')"`
		shellCall(cmdZellijKill)
	},

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zellijKillCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
