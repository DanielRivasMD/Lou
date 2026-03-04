/*
Copyright © 2026 Daniel Rivas <danielrivasmd@gmail.com>

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
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	harvestCmd := MakeCmd("harvest", runHarvest,
		WithArgs(cobra.ExactArgs(1)), // enforce exactly one argument
	)
	rootCmd.AddCommand(harvestCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runHarvest(cmd *cobra.Command, args []string) {
	const op = "lou.harvest"

	// Determine extension (default "sh")
	ext := "sh"
	if len(args) == 1 {
		ext = args[0]
	}

	// Build the command with the chosen extension
	cmdStr := fmt.Sprintf(`fd --extension %s | while read f; do echo "=== $f ==="; cat "$f"; echo; done`, ext)

	horus.CheckErr(
		domovoi.ExecSh(cmdStr),
		horus.WithOp(op),
		horus.WithMessage("failed to harvest files"),
		horus.WithCategory("EXEC_ERROR"),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
