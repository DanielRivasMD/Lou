// cmd/resize.go
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

// flags for resizing
var (
	floatHeight string
	floatWidth  string
	floatX      string
	floatY      string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// zResizeCmd resizes a random floating Zellij pane using percentage flags.
var zResizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Anchor and resize a random floating pane",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Resize one random floating pane to a percentage of screen size and move it to an anchor point.
`,
	Example: `
  lou resize --height 80% --width 70% --x 5% --y 10%
`,
	Run: func(cmd *cobra.Command, args []string) {
		// build shell script, substituting in flag values
		cmdResize := fmt.Sprintf(`
zellij action change-floating-pane-coordinates \
  --pane-id $ZELLIJ_PANE_ID \
  --height %s \
  --width %s \
  --x %s \
  --y %s
		zellij action rename-pane canvas
`, floatHeight, floatWidth, floatX, floatY)
		domovoi.ExecCmd("bash", "-c", cmdResize)

	},
}

func init() {
	rootCmd.AddCommand(zResizeCmd)

	// bind flags with defaults
	zResizeCmd.Flags().StringVarP(&floatHeight, "height", "H", "100%", "pane height as percentage")

	zResizeCmd.Flags().StringVarP(&floatWidth, "width", "W", "95%", "pane width as percentage")

	zResizeCmd.Flags().StringVarP(&floatX, "x", "X", "10", "horizontal offset as percentage")

	zResizeCmd.Flags().StringVarP(&floatY, "y", "Y", "0", "vertical offset as percentage")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
