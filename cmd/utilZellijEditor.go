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

// helpEditorShort returns the one‐line Short description.
func helpEditorShort(editor string) string {
	return fmt.Sprintf("view data in a floating zellij window using %s", editor)
}

// helpEditorLong returns the multi‐line Long description.
func helpEditorLong(editor string) string {
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

// helpEditorExample returns the example usage snippet.
func helpEditorExample(editor string) string {
	return chalk.White.Color("lou") + " " +
		chalk.White.Color(chalk.Bold.TextStyle(editor)) + " " +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>"))
}

// runEditor returns the Run func for a given editor. It
// 1) resolves geometry via resolveLayoutGeometry(layoutFlag)
// 2) builds the zellij command
// 3) execs it via domovoi.ExecCmd
// runEditor returns the cobra‐Run function for your floating editor.
//   - call is the program you actually exec (e.g. "micro", "helix")
//   - editorOverride, if supplied, becomes the --name for the zellij window
func runEditor(call string, editorOverride ...string) func(cmd *cobra.Command, args []string) {
	// determine the name used in `--name`
	editor := call
	if len(editorOverride) > 0 {
		editor = editorOverride[0]
	}

	return func(cmd *cobra.Command, args []string) {
		geom, err := resolveLayoutGeometry(layoutFlag)
		if err != nil {
			// handle error
		}

		// build the zellij invocation
		cmdEditor := fmt.Sprintf(
			`zellij run --name %s --close-on-exit --floating --pinned true \
--height %s \
--width  %s \
--x      %s \
--y      %s \
-- `,
			editor,
			geom.Height, geom.Width, geom.X, geom.Y,
		)
		// append the actual binary you want to run
		cmdEditor += call

		// if the user passed a file (or any single positional), tack it on
		if len(args) > 0 {
			cmdEditor += " " + args[0]
		}

		domovoi.ExecCmd("bash", "-c", cmdEditor)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
