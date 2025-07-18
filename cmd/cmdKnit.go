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
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	knitFile string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// knitCmd
var knitCmd = &cobra.Command{
	Use:   "knit",
	Short: "Compile markdown files using R",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.White.Color("Lou") + ` leverage R to render Markdown files into polished outputs
`,

	Example: chalk.White.Color("lou") + ` ` + chalk.White.Color("knit") + ` ` + chalk.Dim.TextStyle(chalk.Blue.Color("--file")) + ` ` + chalk.Dim.TextStyle("<file>"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// execute command
		cmdKnit := fmt.Sprintf(`R --slave -e "rmarkdown::render('%s')" > /dev/null`, knitFile)
		domovoi.ExecCmd("bash", "-c", cmdKnit)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(knitCmd)

	// flags
	knitCmd.Flags().StringVarP(&knitFile, "file", "f", "", "File to compile")

	knitCmd.MarkFlagRequired("file")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
