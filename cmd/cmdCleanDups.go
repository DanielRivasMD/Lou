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
	"regexp"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
const regParenthesis = `\(\d\)\w*` // backticks here contain expression

// declare regex
var (
	// Define the regular expression pattern for matching file names.
	fileNamePattern = regexp.MustCompile(regParenthesis)

	loc string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// cleanCmd
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"C"},
	Short:   "Clean duplicates",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` clean duplicates at a target location
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("clean") + `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("clean") + ` -l $(pwd)` + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// execute logic
		cleanDir(loc)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(cleanCmd)

	// flags
	cleanCmd.Flags().StringVarP(&loc, "location", "", "/Users/drivas/Downloads/", "Location to clean")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
