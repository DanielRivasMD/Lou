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
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var zfEzaCmd = &cobra.Command{
	Use:   "eza " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<path>")),
	Short: `view data in a floating zellij window using eza`,
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` +
		`view data in a floating ` + chalk.Cyan.Color(chalk.Italic.TextStyle("zellij")) + ` window using ` + chalk.Cyan.Color(chalk.Italic.TextStyle("eza")) + `
`,

	Example: chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("eza")) + ` ` + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<path>")),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		layoutName := layoutFlag

		geom, _ := resolveLayoutGeometry(layoutName)

		cmdEza := fmt.Sprintf(`
		zellij run --name canvas --floating --pinned true \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, geom.Height, geom.Width, geom.X, geom.Y)
		cmdEza += `eza --header --long --icons --classify --git --group --color=always`

		// validate input
		if len(args) > 0 {
			path := args[0]
			cmdEza += " " + path
		}

		domovoi.ExecCmd("bash", "-c", cmdEza)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zfEzaCmd)
	zfCmd.AddCommand(zfEzaCmd)

	zfEzaCmd.Flags().StringVarP(&layoutFlag, "layout", "", "default", "")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
