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

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// knitCmd
var knitCmd = &cobra.Command{

	Use:   "knit",
	Short: "" + chalk.Yellow.Color("cobra") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("tabularasa") + ` help ` + chalk.Yellow.Color("knit"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		shcmds := []string{
			// "R --slave -e "rmarkdown::render('$1')" > /dev/null",
		}

		for _, shcmd := range shcmds {
			err, stdout, stderr := shellCall(shcmd)
			if err != nil {
			  fmt.Println(err.Error())
			}

			lineBreaks()
			fmt.Println()
			fmt.Println(stdout)
			fmt.Println(stderr)

		}

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	osCmd.AddCommand(knitCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////

