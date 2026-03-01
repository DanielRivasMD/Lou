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

func init() {
	knitCmd := MakeCmd("knit", runKnit,
		WithArgs(cobra.ExactArgs(1)),
	)
	rootCmd.AddCommand(knitCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runKnit(cmd *cobra.Command, args []string) {
	const op = "lou.knit"
	inputFile := args[0]
	cmdKnit := fmt.Sprintf(`R --slave -e "rmarkdown::render('%s')" > /dev/null`, inputFile)

	horus.CheckErr(
		domovoi.ExecCmd("bash", "-c", cmdKnit),
		horus.WithOp(op),
		horus.WithMessage("Unable to knit Markdown file"),
		horus.WithCategory("run_error"),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
