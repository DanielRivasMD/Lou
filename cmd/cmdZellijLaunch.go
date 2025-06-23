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

	"github.com/DanielRivasMD/horus"
	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// Global flag variable that holds the target directory, if provided.
var (
	launchTarget string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// zellijLaunchCmd builds and prints the shell command you should run to open
// a new Zellij tab with your custom layout. If --target is set, it wraps the
// launch call in a cd to that directory and back to your original location.
var zellijLaunchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Prepare the shell command to start a new Zellij tab",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
		`

` + chalk.Green.Color("Lou") + ` helps you craft the exact shell invocation
you need to launch a styled Zellij tab. It will print out:

    launch --layout $HOME/.lou/layouts/launch.kdl --name "<dirname>"

Optionally, it can wrap that in 'cd <target> && ... && cd <original>'.`,
	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("launch --target /path/to/project") + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		const cmdZellijLaunch = `zellij action write-chars "zellij --new-session-with-layout $HOME/.config/zellij/layouts/rust.kdl"; zellij action write 13`
		op := "cmd-launch"

		if launchTarget != "" {
			// build a compound command: new-tab + write-chars + ENTER
			const cmdZellijTab = `zellij action new-tab --layout $HOME/.lou/layouts/launch.kdl --name "$( [ "$PWD" = "$HOME" ] && echo "~" || basename "$PWD" )"`
			cmdZellijTabLaunch := cmdZellijTab + "; " + cmdZellijLaunch

			// recall current dir or panic
			originalDir, err := domovoi.RecallDir()
			if err != nil {
				panic(horus.Wrap(err, op, "failed to recall current directory"))
			}
			// ensure we always revert, even on panic
			defer func() {
				if err := domovoi.ChangeDir(originalDir); err != nil {
					panic(horus.Wrap(err, op, "failed to revert to original directory"))
				}
			}()

			// change into target or panic
			if err := domovoi.ChangeDir(launchTarget); err != nil {
				panic(horus.Wrap(
					err, op,
					fmt.Sprintf("failed to change directory to %q", launchTarget),
				))
			}
			// execute the Zellij commands or panic
			if err := domovoi.ExecSh(cmdZellijTabLaunch); err != nil {
				panic(horus.Wrap(err, op, "failed to launch new Zellij session"))
			}
		} else {
			// no target → just send the write-chars + ENTER sequence
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

	// Bind the --target / -t flag to targetDir.
	zellijLaunchCmd.Flags().
		StringVarP(&launchTarget, "target", "t", "", "If set, cd into this path before printing the launch command (and return afterward)")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
