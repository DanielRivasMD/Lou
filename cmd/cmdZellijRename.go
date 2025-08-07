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
	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var zRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: `rename current tab with working directory`,
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` +
		`rename the current ` + chalk.Cyan.Color(chalk.Italic.TextStyle("zellij")) + ` tab effortlessly using the working directory
`,

	Example: chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("rename")),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		cmdRename := `zellij action rename-tab "$( [ "$PWD" = "$HOME" ] && echo "~" || basename "$PWD" )"`
		domovoi.ExecCmd("bash", "-c", cmdRename)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zRenameCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
