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

// buildCmd
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: `Build bibliography from ` + chalk.Cyan.Color("nbib") + ` files.`,
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` will relocate ` + chalk.Cyan.Color("nbib") + ` files from a target location to a target destination. If target destination does not exist, it will be created.

Next, ` + chalk.Cyan.Color("nbib") + ` files will be compiled into a ` + chalk.Cyan.Color("LaTeX") + ` compatible reference library.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// execute logic
		fmt.Println("build called")
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	bibloCmd.AddCommand(buildCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
