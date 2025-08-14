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

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var hiddenCmd = &cobra.Command{
	Use:   "hidden " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("[off|on]")),
	Short: "Toggle Finder visibility of hidden files",

	// TODO: update by using functions
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` control the visibility of hidden files in Finder on macOS
Toggle between showing and hiding hidden files using the appropriate argument
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("finder") + ` off` + `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("finder") + ` on` + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"off", "on"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: runHidden,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(hiddenCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runHidden(cmd *cobra.Command, args []string) {
	const op = "lou.hidden"
	var cmdFinder string
	arg := args[0]

	switch arg {
	case "off":
		cmdFinder = `defaults write com.apple.Finder AppleShowAllFiles false && killall Finder`
	case "on":
		cmdFinder = `defaults write com.apple.Finder AppleShowAllFiles true && killall Finder`
	default:
		fmt.Printf("Invalid argument: %s\n", arg)
	}

	horus.CheckErr(
		domovoi.ExecCmd("bash", "-c", cmdFinder),
		horus.WithOp(op),
		horus.WithMessage("Unable to toogle hidden files in Finder"),
		horus.WithCategory("run_error"),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
