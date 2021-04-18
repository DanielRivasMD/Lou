/*
Copyright © 2021 Daniel Rivas <danielrivasmd@gmail.com>

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

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/DanielRivasMD/Lou/auxiliary"
	"github.com/atrox/homedir"
	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

// annotateCmd represents the annotate command
var annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Lou relocates files for annoatation",
	Long: `Lou will relocate pdf file for annotation
	while keeping a reference to original location.

	When overwritten, annotated file will remain intact.`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		// find home directory
		home, errHomedir := homedir.Dir()
		if errHomedir != nil {
			fmt.Println(errHomedir)
			os.Exit(1)
		}

		//command line flags
		target, _ := cmd.Flags().GetString("target")

		// lineBreaks
		auxiliary.LineBreaks()

		// buffers
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		// shell call
		commd := home + "/Factorem/Lou/sh/note.sh"
		shCmd := exec.Command(commd, target)

		// run
		shCmd.Stdout = &stdout
		shCmd.Stderr = &stderr
		_ = shCmd.Run()

		// stdout
		fmt.Println(chalk.Cyan.Color(stdout.String()))

		// stderr
		if stderr.String() != "" {
			fmt.Println(chalk.Red.Color(stderr.String()))
		}

		// lineBreaks
		auxiliary.LineBreaks()

	},

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

func init() {
	bibloCmd.AddCommand(annotateCmd)

	////////////////////////////////////////////////////////////////////////////////////////////////////

	// flags
	annotateCmd.Flags().StringP("target", "t", "", "article to annoate")
	annotateCmd.MarkFlagRequired("target")

	////////////////////////////////////////////////////////////////////////////////////////////////////

}
