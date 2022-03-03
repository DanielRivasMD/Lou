/*
Copyright © 2020 Daniel Rivas <danielrivasmd@gmail.com>

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

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// bibloCmd represents the article command
var bibloCmd = &cobra.Command{
	Use:     "biblo",
	Aliases: []string{"b"},
	Short:   "Handles all biblography operations.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` handles all biblography operations.

For example:
- Reformat articles and their references downloaded by Kopernico.
- Relocate articles manually renamed.
- Relocate theses to the proper archive.
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` biblo format
` + chalk.Cyan.Color("lou") + ` biblo thesis`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"format", "thesis"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// execute logic
		bibloArgs(findHome(), args)

		// clean after relocating
		// determine location
		location, _ := cmd.Flags().GetString("location")

		// execute logic
		cleanDir(location)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(bibloCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func bibloArgs(home string, args []string) {

	// allocate command
	var commd string

	// shell call
	switch args[0] {

	case "format":
		commd = home + "/Factorem/Lou/sh/format.sh"
		runSh(commd)

		commd = home + "/Factorem/Lou/sh/locate.sh"
		runSh(commd)

	case "thesis":
		commd = home + "/Factorem/Lou/sh/thesis.sh"
		runSh(commd)

	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runSh(commd string) {

	// lineBreaks
	lineBreaks()

	// buffers
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// run command
	shCmd := exec.Command(commd)
	shCmd.Stdout = &stdout
	shCmd.Stderr = &stderr
	_ = shCmd.Run()

	// stdout
	fmt.Println(chalk.Cyan.Color(stdout.String()))

	// stderr
	if stderr.String() != "" {
		fmt.Println(chalk.Red.Color(stderr.String()))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
