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

// countCmd
var countCmd = &cobra.Command{

	Use:   "count",
	Short: "" + chalk.Yellow.Color("cobra") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("tabularasa") + ` help ` + chalk.Yellow.Color("count"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"dir", "file"},
	Args:      cobra.ExactValidArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// parse flags
		ϙ_hidden, ε := κ.Flags().GetBool("hidden")
		checkErr(ε)

		hidden := ""
		if ϙ_hidden {
			hidden = "--hidden "
		}

		// parse flags
		ϙ_no_ignore, ε := κ.Flags().GetBool("no-ignore")
		checkErr(ε)

		no_ignore := ""
		if ϙ_no_ignore {
			no_ignore = "--no-ignore "
		}

		// declare cmd
		cmd := "fd . "

		switch args[0] {
		case "dir":
			shcmd := cmd + hidden + no_ignore + "--type=d --max-depth=1 | /usr/bin/wc -l"
			ε, stdout, _ := shellCall(shcmd)
			checkErr(ε)
			fmt.Print(chalk.Yellow.Color("Number of dirs: "), stdout)
		case "file":
			shcmd := cmd + hidden + no_ignore + "--type=f --max-depth=1 | /usr/bin/wc -l"
			ε, stdout, _ := shellCall(shcmd)
			checkErr(ε)
			fmt.Print(chalk.Yellow.Color("Number of files: "), stdout)
		}

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	osCmd.AddCommand(countCmd)

	// flags
	countCmd.Flags().BoolP("hidden", "H", false, "account hidden files/dirs.")
	countCmd.Flags().BoolP("no-ignore", "I", false, "don't respect ignore files.")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
