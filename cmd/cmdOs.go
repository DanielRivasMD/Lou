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

import (
	"fmt"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

func pathPrint(envvar string) {

	// 1. Read $PATH environment variable
	echo := `echo $` + envvar
	sout, _, err := domovoi.CaptureExecCmd("zsh", "-c", echo)
	horus.CheckErr(err)

	// 2. Split on colon
	segments := strings.Split(sout, ":")

	// 3. Print
	domovoi.LineBreaks(true)
	for s := range segments {
		if s == len(segments)-1 {
			fmt.Printf(segments[s])
		} else {
			fmt.Println(segments[s])
		}
	}
	domovoi.LineBreaks(true)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// osCmd
var osCmd = &cobra.Command{
	Use:   "os",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("") + ` help ` + chalk.Yellow.Color("") + chalk.Yellow.Color("os"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(osCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
