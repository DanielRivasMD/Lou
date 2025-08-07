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
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var knitCmd = &cobra.Command{
	Use:   "knit " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),
	Short: "compile a Markdown file using R",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + `leverage ` + chalk.Cyan.Color(chalk.Italic.TextStyle("R")) + ` to render a Markdown file into a polished output
`,
	Example: chalk.White.Color("lou") + " " + chalk.Bold.TextStyle(chalk.White.Color("knit")) + " " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Args: cobra.ExactArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		inputFile := args[0]
		cmdKnit := fmt.Sprintf(`R --slave -e "rmarkdown::render('%s')" > /dev/null`, inputFile)
		err := domovoi.ExecCmd("bash", "-c", cmdKnit)
		horus.CheckErr(err)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(knitCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
