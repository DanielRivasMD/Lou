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

var zfFloatCmd = &cobra.Command{
	Use:   "float",
	Short: `launch floating zellij window with ease`,
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` +
		`launch a floating ` + chalk.Cyan.Color(chalk.Italic.TextStyle("zellij")) + ` window using ` + chalk.Cyan.Color(chalk.Italic.TextStyle("float")) + ` with ease
`,

	Example: chalk.White.Color("lou") + ` ` + chalk.White.Color(chalk.Bold.TextStyle("float")),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Args:      cobra.MaximumNArgs(1),
	ValidArgs: validLayouts,

	Run: runFloat,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zfFloatCmd)
	zfCmd.AddCommand(zfFloatCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runFloat(cmd *cobra.Command, args []string) {

	layoutName := "default"
	if len(args) == 1 {
		layoutName = args[0]
	}

	geom, _ := resolveLayoutGeometry(layoutName)

	cmdFloat := fmt.Sprintf(`
		zellij run --name canvas --close-on-exit --floating \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, geom.Height, geom.Width, geom.X, geom.Y)
	cmdFloat += `zsh`
	domovoi.ExecCmd("bash", "-c", cmdFloat)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
