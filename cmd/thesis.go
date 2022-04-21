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

// thesisCmd represents the thesis command
var thesisCmd = &cobra.Command{
	Use:   "thesis",
	Short: "Collect theses.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` will relocate theses.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// execute logic
		fmt.Println("thesis called")

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	bibloCmd.AddCommand(thesisCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////
