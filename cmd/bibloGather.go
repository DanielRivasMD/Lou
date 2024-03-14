/*
Copyright © 2022 Daniel Rivas <danielrivasmd@gmail.com>

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

// gatherCmd
var gatherCmd = &cobra.Command{
	Use:   "gather",
	Short: "Collect all articles & references.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` will relocate articles (` + chalk.Cyan.Color("pdf") + `) and their references (` + chalk.Cyan.Color("nbib") + `) downloaded by ` + chalk.Yellow.Color("Kopernico") + `.

Additionally, ` + chalk.Green.Color("Lou") + ` will also relocate manually renamed articles.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// execute logic
		bibloArgs(findHome())

		fmt.Println()

		// determine location
		location, _ := κ.Flags().GetString("location")

		// clean after relocating
		cleanDir(location)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	bibloCmd.AddCommand(gatherCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
