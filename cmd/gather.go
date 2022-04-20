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

// gatherCmd represents the gather command
var gatherCmd = &cobra.Command{
	Use:   "gather",
	Short: "Collect all articles & references.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` will relocate articles (` + chalk.Cyan.Color("pdf") + `) and their references (` + chalk.Cyan.Color("nbib") + `) downloaded by ` + chalk.Yellow.Color("Kopernico") + `.

Additionally, ` + chalk.Green.Color("Lou") + ` will also relocate manually renamed articles.
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gather called")
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	bibloCmd.AddCommand(gatherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
