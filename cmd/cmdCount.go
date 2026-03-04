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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	hidden   bool
	noIgnore bool
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	countCmd := MakeCmd("count", runCount,
		WithArgs(cobra.ExactArgs(1)), // enforce exactly one argument
		WithValidArgs([]string{"dir", "file"}),
	)
	rootCmd.AddCommand(countCmd)

	countCmd.Flags().BoolVarP(&hidden, "hidden", "", false, "include hidden files and dirs")
	countCmd.Flags().BoolVarP(&noIgnore, "no-ignore", "", false, "do not respect ignore config")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCount(cmd *cobra.Command, args []string) {
	const op = "lou.count"

	target := args[0]

	// Build the fd call
	cmdStr := "fd ."
	if hidden {
		cmdStr += " --hidden"
	}
	if noIgnore {
		cmdStr += " --no-ignore"
	}

	switch target {
	case "dir":
		cmdStr += " --type=d --max-depth=1 | wc -l"
	case "file":
		cmdStr += " --type=f --max-depth=1 | wc -l"
	default:
		horus.CheckErr(
			fmt.Errorf("unsupported mode %q; use \"dir\" or \"file\"", target),
			horus.WithOp(op),
			horus.WithCategory("USAGE_ERROR"),
			horus.WithExitCode(2),
		)
	}

	if err := domovoi.ExecSh(cmdStr); err != nil {
		herr := horus.PropagateErr(
			op,
			"SYS_CMD",
			"failed to execute count command",
			err,
			map[string]any{
				"target":    target,
				"hidden":    hidden,
				"no_ignore": noIgnore,
				"command":   cmdStr,
			},
		)
		horus.CheckErr(herr, horus.WithFormatter(horus.SimpleColoredFormatter))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
