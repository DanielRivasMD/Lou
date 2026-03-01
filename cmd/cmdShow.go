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

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	listFile string
)

type Function struct {
	Shell       string
	Name        string
	Description string
	Arguments   string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	showCmd := MakeCmd("show", runShow)
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringVarP(&listFile, "file", "f", "", "File to review")
	horus.CheckErr(showCmd.MarkFlagRequired("file"))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runShow(cmd *cobra.Command, args []string) {
	const op = "lou.show"

	// collect documentation
	functions, err := parseFile(listFile)
	horus.CheckErr(
		err,
		horus.WithOp(op),
		horus.WithCategory("PARSE_ERROR"),
		horus.WithMessage("failed to parse shell functions"),
	)

	// generate & print Markdown
	markdown := generateMD(functions)
	fmt.Println(markdown)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
