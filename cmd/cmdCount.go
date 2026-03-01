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
	"os"

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
		WithArgs(cobra.MaximumNArgs(1)),
		WithValidArgs([]string{"dir", "file"}),
	)
	rootCmd.AddCommand(countCmd)

	countCmd.Flags().BoolVarP(&hidden, "hidden", "", false, "include hidden files and dirs")
	countCmd.Flags().BoolVarP(&noIgnore, "no-ignore", "", false, "do not respect ignore config")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCount(cmd *cobra.Command, args []string) {
	const op = "cmd.count"

	// 1) validate args
	if len(args) != 1 {
		herr := horus.NewCategorizedHerror(
			op,
			"USAGE_ERROR",
			fmt.Sprintf("accepts 1 arg(s), received %d", len(args)),
			nil,
			nil,
		)
		he, _ := horus.AsHerror(herr)
		fmt.Fprintln(cmd.ErrOrStderr(), horus.SimpleColoredFormatter(he))
		_ = cmd.Usage()
		os.Exit(1)
	}

	target := args[0]

	// 2) build the fd invocation
	fdCmd := "fd ."
	if hidden {
		fdCmd += " --hidden"
	}
	if noIgnore {
		fdCmd += " --no-ignore"
	}

	switch target {
	case "dir":
		fdCmd += " --type=d --max-depth=1 | wc -l"
	case "file":
		fdCmd += " --type=f --max-depth=1 | wc -l"
	default:
		herr := horus.NewCategorizedHerror(
			op,
			"USAGE_ERROR",
			fmt.Sprintf("unsupported mode %q; use \"dir\" or \"file\"", target),
			nil,
			nil,
		)
		he, _ := horus.AsHerror(herr)
		fmt.Fprintln(cmd.ErrOrStderr(), horus.SimpleColoredFormatter(he))
		_ = cmd.Usage()
		os.Exit(1)
	}

	// 3) execute the command and fatal-exit on error
	if err := domovoi.ExecSh(fdCmd); err != nil {
		// wrap & add context
		herr := horus.PropagateErr(
			op,
			"SYS_CMD",
			"failed to execute count command",
			err,
			map[string]any{
				"target":    target,
				"hidden":    hidden,
				"no_ignore": noIgnore,
				"command":   fdCmd,
			},
		)
		horus.CheckErr(
			herr,
			horus.WithFormatter(horus.SimpleColoredFormatter),
		)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
