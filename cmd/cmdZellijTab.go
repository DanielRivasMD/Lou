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

// Global flag variables
var (
	tabTarget  string
	layoutType string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: consider refactor as subcommands
// zellijTabCmd launches a new Zellij tab with a custom layout.
// Optionally, it can switch to a specified directory before initiating the new tab,
// and then revert to the original directory.
// The --layout (or -l) flag selects between "tab" (default), "explore", or "repl".
var zellijTabCmd = &cobra.Command{
	Use:   "tab",
	Short: "Launch a new Zellij tab with style",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
		`

` + chalk.Green.Color("Lou") + ` allows you to launch a new tab in the active Zellij session effortlessly,
using a custom layout.

Optionally, you can change to a specified directory prior to launching the tab,
and then return to your original directory.

Available layouts: "tab", "explore", "repl"`,
	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("tab --target /path/to/dir --layout explore") + `
`,

	Run: func(cmd *cobra.Command, args []string) {

		const op = "cmd-tab"

		// Validate layout using a known set
		validLayouts := map[string]bool{
			"tab":     true,
			"explore": true,
			"repl":    true,
		}

		if !validLayouts[layoutType] {
			err := fmt.Errorf("invalid layout %q: must be one of [tab, explore, repl]", layoutType)
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("unsupported layout provided"))
			return // gracefully exit after logging the error
		}

		// Validate layout
		switch layoutType {
		case "tab", "explore", "repl":
			// valid
		default:
			layoutType = "tab"
		}

		// Build tab launch command
		cmdZellijTab := fmt.Sprintf(
			`zellij action new-tab --layout $HOME/.lou/layouts/%s.kdl --name $( [ $PWD = $HOME ] && echo \"~\" || basename $PWD )`,
			layoutType,
		)

		if tabTarget != "" {
			originalDir, err := domovoi.RecallDir()
			if err != nil {
				horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to recall current directory"))
			}

			if err := domovoi.ChangeDir(tabTarget); err != nil {
				domovoi.ChangeDir(originalDir)
				horus.CheckErr(err, horus.WithOp(op), horus.WithMessage(fmt.Sprintf("failed to change directory to %q", tabTarget)))
			}

			if err := domovoi.ExecSh(cmdZellijTab); err != nil {
				domovoi.ChangeDir(originalDir)
				horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to launch new tab"))
			}
		} else {
			if err := domovoi.ExecSh(cmdZellijTab); err != nil {
				horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to launch new tab"))
			}
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zellijTabCmd)

	// Bind flags
	zellijTabCmd.Flags().StringVarP(&tabTarget, "target", "t", "", "If provided, change to this directory before launching the new tab and then revert to the original directory")
	zellijTabCmd.Flags().StringVarP(&layoutType, "layout", "l", "tab", "Layout to use: tab (default), explore, or repl")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
