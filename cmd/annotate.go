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
	"os/exec"

	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

// annotateCmd represents the annotate command
var annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Relocates files for annotation",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` will relocate pdf file for annotation
	while keeping a reference to original location.

	When overwritten, annotated file will remain intact.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		//command line flags
		target, _ := cmd.Flags().GetString("target")

		// execute logic
		annotateFile(findHome(), target)

	},
}





func annotateFile(home, target string) {

	// lineBreaks
	lineBreaks()

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// buffers
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// shell call
	commd := home + "/Factorem/Lou/sh/note.sh"
	shCmd := exec.Command(commd, target)

func init() {
	bibloCmd.AddCommand(annotateCmd)
	// run
	shCmd.Stdout = &stdout
	shCmd.Stderr = &stderr
	_ = shCmd.Run()

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// stdout
	fmt.Println(chalk.Cyan.Color(stdout.String()))

	// flags
	annotateCmd.Flags().StringP("target", "t", "", "article to annoate")
	annotateCmd.MarkFlagRequired("target")
	// stderr
	if stderr.String() != "" {
		fmt.Println(chalk.Red.Color(stderr.String()))
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////
	// lineBreaks
	lineBreaks()

}
