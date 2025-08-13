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

// helpShort returns the one‐line Short description.
func helpShort(editor string) string {
	return fmt.Sprintf("view data in a floating zellij window using %s", editor)
}

// helpLong returns the multi‐line Long description.
func helpLong(editor string) string {
	header := chalk.Green.Color(
		chalk.Bold.TextStyle("Daniel Rivas "),
	) +
		chalk.Dim.TextStyle(
			chalk.Italic.TextStyle("<danielrivasmd@gmail.com>"),
		)

	body := fmt.Sprintf(
		"\n\nview data in a floating %szellij%s window using %s",
		chalk.Cyan.Color(""),
		chalk.Cyan.Color(""),
		chalk.Cyan.Color(editor),
	)

	return header + chalk.Dim.TextStyle(body)
}

// helpExample returns the example usage snippet.
func helpExample(editor string) string {
	return chalk.White.Color("lou") + " " +
		chalk.White.Color(chalk.Bold.TextStyle(editor)) + " " +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>"))
}

// runEditor returns the Run func for a given editor. It
// 1) resolves geometry via resolveLayoutGeometry(layoutFlag)
// 2) builds the zellij command
// 3) execs it via domovoi.ExecCmd
func runEditor(editor string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		layoutName := layoutFlag
		geom, _ := resolveLayoutGeometry(layoutName)

		cmdEditor := fmt.Sprintf(`
		zellij run --name canvas --close-on-exit --floating --pinned true \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, geom.Height, geom.Width, geom.X, geom.Y)
		cmdEditor += editor

		if len(args) > 0 {
			file := args[0]
			cmdEditor += " " + file
		}

		fmt.Println(cmdEditor)

		domovoi.ExecCmd("bash", "-c", cmdEditor)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
