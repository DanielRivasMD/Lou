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
	"os/exec"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	amassFlags struct {
		copy bool
		list bool
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func AmassCmd() *cobra.Command {
	d := horus.Must(domovoi.GlobalDocs())
	cmd := horus.Must(d.MakeCmd("amass", runAmass,
		domovoi.WithArgs(cobra.MinimumNArgs(1)),
	))

	cmd.Flags().BoolVarP(&amassFlags.copy, "copy", "c", false, "copy output to clipboard (pbcopy on macOS, xclip/xsel on Linux)")
	cmd.Flags().BoolVarP(&amassFlags.list, "list", "l", false, "list only file names")

	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runAmass(cmd *cobra.Command, args []string) {
	const op = "lou.amass"

	fdArgs := []string{}
	for _, ext := range args {
		fdArgs = append(fdArgs, "--extension", ext)
	}

	var output []byte
	var err error

	if amassFlags.list {
		output, err = exec.Command("fd", fdArgs...).Output()
		if err != nil {
			horus.CheckErr(err,
				horus.WithOp(op),
				horus.WithMessage("failed to list files"),
				horus.WithCategory("EXEC_ERROR"),
			)
		}
	} else {
		fdCmdStr := "fd"
		for _, ext := range args {
			fdCmdStr += " --extension " + ext
		}
		fullCmd := fdCmdStr + ` | while read f; do echo "=== $f ==="; cat "$f"; echo; done`

		output, err = exec.Command("bash", "-c", fullCmd).Output()
		if err != nil {
			if exitErr, ok := err.(*exec.ExitError); ok {
				horus.CheckErr(fmt.Errorf("%s: %s", err, exitErr.Stderr),
					horus.WithOp(op),
					horus.WithMessage("failed to amass files"),
					horus.WithCategory("EXEC_ERROR"),
				)
			} else {
				horus.CheckErr(err,
					horus.WithOp(op),
					horus.WithMessage("failed to amass files"),
					horus.WithCategory("EXEC_ERROR"),
				)
			}
		}
	}

	if amassFlags.copy {
		if err := copyToClipboard(output); err != nil {
			horus.CheckErr(err,
				horus.WithOp(op),
				horus.WithMessage("failed to copy to clipboard"),
				horus.WithCategory("EXEC_ERROR"),
			)
		}
	} else {
		fmt.Print(string(output))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
