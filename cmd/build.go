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

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "// TODO: building",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` relocates ` + chalk.Cyan.Color("nbib") + ` files from a target location to a target destination. If target destination does not exist, it will be created.

Next, ` + chalk.Cyan.Color("nbib") + ` files will be compiled into a LaTeX compatible reference library.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	bibloCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// buildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
