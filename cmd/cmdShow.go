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

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	listFile string
)

// Function represents a parsed shell function
type Function struct {
	Shell       string
	Name        string
	Description string
	Arguments   string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// listShCmd
var listShCmd = &cobra.Command{
	Use:     "show --file <path>",
	Aliases: []string{"s"},
	Short:   "List shell functions from a specified file",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` extract and list shell function details, including descriptions and arguments, from a specified file
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("show") + ` --file functions.sh` + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// collect documentation
		functions, ε := parseFile(listFile)
		checkErr(ε)

		// generate & print Markdown
		markdown := generateMD(functions)
		fmt.Println(markdown)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(listShCmd)

	// flags
	listShCmd.Flags().StringVarP(&listFile, "file", "f", "", "File to review")

	listShCmd.MarkFlagRequired("file")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
