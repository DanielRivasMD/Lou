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
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var zResizeCmd = &cobra.Command{
	Use:   "resize",
	Short: `anchor and resize a random floating pane`,
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` +
		`resize one random floating pane to a percentage of screen size and move it to an anchor point
`,

	Example: chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("resize")) + ` ` +
		chalk.White.Color(chalk.Italic.TextStyle("--height")) + ` ` + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<100%>")) + ` ` +
		chalk.White.Color(chalk.Italic.TextStyle("--width")) + ` ` + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<95%>")) + ` ` +
		chalk.White.Color(chalk.Italic.TextStyle("--x")) + ` ` + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<10>")) + ` ` +
		chalk.White.Color(chalk.Italic.TextStyle("--y")) + ` ` + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<0>")) + `
` +
		// Preset examples
		chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("resize")) + ` full` + `
` +
		chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("resize")) + ` half-left` + `
` +
		chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("resize")) + ` half-right`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Args:      cobra.MaximumNArgs(1),
	ValidArgs: []string{"full", "half-left", "half-right", "top-left", "bottom-left", "top-right", "bottom-right"},

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			switch args[0] {
			case "full":
				floatHeight, floatWidth, floatX, floatY = "100%", "100%", "0", "0"
			case "half-left":
				floatHeight, floatWidth, floatX, floatY = "100%", "50%", "0", "0"
			case "half-right":
				floatHeight, floatWidth, floatX, floatY = "100%", "50%", "50%", "0"
			case "top-left":
				floatHeight, floatWidth, floatX, floatY = "45%", "45%", "0", "0"
			case "bottom-left":
				floatHeight, floatWidth, floatX, floatY = "45%", "45%", "0", "60%"
			case "top-right":
				floatHeight, floatWidth, floatX, floatY = "45%", "45%", "60%", "0"
			case "bottom-right":
				floatHeight, floatWidth, floatX, floatY = "45%", "45%", "60%", "60%"
			default:
				herr := horus.NewCategorizedHerror(
					"resize",
					"validation",
					"invalid preset: expected one of [full, half-left, half-right]",
					nil,
					map[string]any{"got": args[0]},
				)
				horus.CheckErr(herr)
				return
			}
		}

		cmdResize := fmt.Sprintf(`
		zellij action rename-pane canvas
		zellij action change-floating-pane-coordinates --pane-id $ZELLIJ_PANE_ID \
		--height %s \
		--width %s \
		--x %s \
		--y %s
	`, floatHeight, floatWidth, floatX, floatY)
		domovoi.ExecCmd("bash", "-c", cmdResize)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zResizeCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
