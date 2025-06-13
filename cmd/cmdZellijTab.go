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
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// Global flag variable that holds the target directory, if provided.
var (
	targetDir string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// zellijTabCmd launches a new Zellij tab with a custom layout. Optionally, it can switch
// to a specified directory before initiating the new tab, and then revert to the original directory.
// If the --target (or -t) flag is not provided, the new tab is launched from the current directory.
var zellijTabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Launch a new Zellij tab with style",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
		`

` + chalk.Green.Color("Lou") + ` allows you to launch a new tab in the active Zellij session effortlessly,
using a custom layout.

Optionally, you can change to a specified directory prior to launching the tab,
and then return to your original directory.`,
	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("tab --target /path/to/directory") + `
`,
	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// The command string to launch a new Zellij tab.
		// It sets the tab name as either "~" (if in $HOME) or the basename of the current directory.
		const cmdZellijTab = `zellij action new-tab --layout $HOME/.lou/layouts/tab.kdl --name "$( [ "$PWD" = "$HOME" ] && echo "~" || basename "$PWD" )"`

		// If the targetDir flag is provided, change the directory accordingly.
		if targetDir != "" {
			// Recall the current directory so we can revert back later.
			originalDir, err := domovoi.RecallDir()
			if err != nil {
				fmt.Println(chalk.Red.Color("Error recalling current directory: " + err.Error()))
				return
			}

			// Change to the target directory before launching the new tab.
			err = domovoi.ChangeDir(targetDir)
			if err != nil {
				fmt.Println(chalk.Red.Color("Error changing directory to target: " + err.Error()))
				return
			}

			// Launch the new tab from the new (target) directory.
			if err := domovoi.ExecCmd("bash", "-c", cmdZellijTab); err != nil {
				fmt.Println(chalk.Red.Color("Error launching new tab: " + err.Error()))
			}

			// Revert back to the original directory.
			if err := domovoi.ChangeDir(originalDir); err != nil {
				fmt.Println(chalk.Red.Color("Error reverting to original directory: " + err.Error()))
			}
		} else {
			// If no target directory is provided, launch the tab from the current directory.
			if err := domovoi.ExecCmd("bash", "-c", cmdZellijTab); err != nil {
				fmt.Println(chalk.Red.Color("Error launching new tab: " + err.Error()))
			}
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zellijTabCmd)

	// Bind the --target (or -t) flag to the targetDir variable.
	zellijTabCmd.Flags().StringVarP(&targetDir, "target", "t", "", "If provided, change to this directory before launching the new tab and then revert to the original directory")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
