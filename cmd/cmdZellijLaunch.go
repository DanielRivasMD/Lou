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
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// Global flag variable
var (
	layoutFile   string
	launchTarget string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// zellijLaunchCmd builds and prints the shell command you should run to open
// a new Zellij tab with your custom layout. The --layout flag is now required.
var zellijLaunchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Prepare the shell command to start a new Zellij tab",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
		`

` + chalk.Green.Color("Lou") + ` helps you craft the exact shell invocation
you need to launch a styled Zellij tab. You must now specify --layout <file.kdl>.`,
	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("launch --layout rust.kdl") + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		op := "cmd-launch"

		// Crash if no layout file was provided
		if layoutFile == "" {
			panic(horus.NewCategorizedHerror(
				op,
				"missing_flag",
				"the --layout flag is required but was not provided",
				nil,
				map[string]any{"flag": "layout"},
			))
		}

		// Build the dynamic write-chars command for creating a new session
		cmdZellijLaunch := fmt.Sprintf(
			`zellij action write-chars "zellij --new-session-with-layout $HOME/.config/zellij/layouts/%s"; zellij action write 13`,
			layoutFile,
		)

		if launchTarget != "" {
			// Build the compound new-tab + write-chars + ENTER command
			const cmdZellijTab = `zellij action new-tab \
--layout $HOME/.lou/layouts/launch.kdl \
--name "$( [ "$PWD" = "$HOME" ] && echo "~" || basename "$PWD" )"`
			fullCmd := cmdZellijTab + "; " + cmdZellijLaunch

			// Recall original dir or panic
			originalDir, err := domovoi.RecallDir()
			if err != nil {
				panic(horus.Wrap(err, op, "failed to recall original directory"))
			}

			// Ensure we always revert, even on panic
			defer func() {
				if err := domovoi.ChangeDir(originalDir); err != nil {
					panic(horus.Wrap(err, op, "failed to revert to original directory"))
				}
			}()

			// Change into target or panic
			if err := domovoi.ChangeDir(launchTarget); err != nil {
				panic(horus.Wrap(
					err, op,
					fmt.Sprintf("failed to change directory to %q", launchTarget),
				))
			}

			// Execute the Zellij commands or panic
			if err := domovoi.ExecSh(fullCmd); err != nil {
				panic(horus.Wrap(err, op, "failed to launch new Zellij session"))
			}
		} else {
			// No target → just send the write-chars + ENTER sequence
			if err := domovoi.ExecSh(cmdZellijLaunch); err != nil {
				panic(horus.Wrap(err, op, "failed to write-chars for new Zellij session"))
			}
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zellijLaunchCmd)

	zellijLaunchCmd.Flags().StringVarP(&layoutFile, "layout", "l", "", "the .kdl layout file to launch (required)")
	zellijLaunchCmd.Flags().StringVarP(&launchTarget, "target", "t", "", "if set, cd into this path before printing the launch command (and return afterward)")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
