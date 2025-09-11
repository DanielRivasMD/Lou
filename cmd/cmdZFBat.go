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

var zfBatCmd = &cobra.Command{
	Use:     "bat " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),
	Short:   "View data in a floating zellij window using bat",
	Long:    helpZFBat,
	Example: exampleZFBat,

	Run: runBat,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zfBatCmd)
	zfCmd.AddCommand(zfBatCmd)

	zfBatCmd.Flags().StringVarP(&layoutFlag, "layout", "", "default", "")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runBat(cmd *cobra.Command, args []string) {
	op := "lou.bat"
	if len(args) < 1 {
		horus.CheckErr(
			horus.NewHerrorErrorf(
				op,
				"bat command requires a file argument",
			),
		)
	}
	file := args[0]

	layoutName := layoutFlag

	geom, _ := resolveLayoutGeometry(layoutName)

	cmdBat := fmt.Sprintf(`
		zellij run --name bat --close-on-exit --floating --pinned true \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, geom.Height, geom.Width, geom.X, geom.Y)
	cmdBat += `bat --paging always`
	cmdBat += " " + file
	// cmdBat += "; tput setaf 5 bold; read -p 'Press any key to continue'"
	domovoi.ExecCmd("bash", "-c", cmdBat)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
