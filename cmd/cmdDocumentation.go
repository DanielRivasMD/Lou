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
	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func DocumentationCmd() *cobra.Command {
	d := horus.Must(domovoi.GlobalDocs())
	return horus.Must(d.MakeCmd("documentation", runDocumentation,
		domovoi.WithArgs(cobra.ExactArgs(1)),
	))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runDocumentation(cmd *cobra.Command, args []string) {
	const op = "lou.documentation"
	arg := args[0]

	cmdDocumentation := `
# grab the list under “Available Commands:”
cmds=( $(
  ` + arg + ` help | 
  sed -n '/Available Commands:/,/Flags:/p' |
  grep -E '^[[:space:]]+[a-z]' | 
  awk '{print $1}'
) )

for sub in "${cmds[@]}"; do
  printf "==================================================\n"
  printf "===== %s \n" "$sub"
  printf "==================================================\n"
  ` + arg + ` help "$sub"
done
`

	domovoi.LineBreaks(true)
	horus.CheckErr(
		domovoi.ExecCmd("zsh", "-c", cmdDocumentation),
		horus.WithOp(op),
		horus.WithMessage("Unable to execute documentation command"),
		horus.WithCategory("run_error"),
	)
	domovoi.LineBreaks(true)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
