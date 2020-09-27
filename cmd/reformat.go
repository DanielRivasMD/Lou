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
	"log"

	"github.com/DanielRivasMD/Lou/lineBreaks"
	"github.com/DanielRivasMD/Lou/shellCall"

	"github.com/labstack/gommon/color"

	"github.com/spf13/cobra"
)

// reformatCmd represents the reformat command
var reformatCmd = &cobra.Command{
	Use:   "reformat",
	Short: "Reformat and relocate",
	Long:  `Article Reform reformats articles, pdf format, and references, ris format, downloaded via Kopernico`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		// lineBreaks
		lineBreaks.LineBreaks()

		// shellCall
		err, out, errout := shellCall.ShellCall("/Users/drivas/Factorem/Lou/sh/reformat.sh")
		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// stdout
		color.Println(color.Cyan(out, color.B))

		// stderr
		if errout != "" {
			color.Println(color.Red(errout, color.B))
		}

		// lineBreaks
		lineBreaks.LineBreaks()

	},

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

func init() {
	articleCmd.AddCommand(reformatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reformatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reformatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
