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

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var zfMdcatCmd = &cobra.Command{
	Use:   "mdcat",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("") + ` help ` + chalk.Yellow.Color("") + chalk.Yellow.Color("mdcat"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			horus.CheckErr(
				horus.NewHerrorErrorf(
					"mdcat",
					"mdcat command requires a file argument",
				),
			)
		}
		file := args[0]

		layoutName := layoutFlag

		geom, _ := resolveLayoutGeometry(layoutName)

		cmdMdcat := fmt.Sprintf(`
		zellij run --name canvas --close-on-exit --floating --pinned true \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, geom.Height, geom.Width, geom.X, geom.Y)
		cmdMdcat += `mdcat --paging always`
		cmdMdcat += " " + file
		domovoi.ExecCmd("bash", "-c", cmdMdcat)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zfMdcatCmd)
	zfCmd.AddCommand(zfMdcatCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
