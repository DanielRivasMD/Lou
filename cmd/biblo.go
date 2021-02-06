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
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/DanielRivasMD/Lou/aux"
	"github.com/mitchellh/go-homedir"
	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

var reformat bool

// bibloCmd represents the article command
var bibloCmd = &cobra.Command{
	Use:     "biblo",
	Aliases: []string{"b"},
	Short:   "Lou handles all biblography operations",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

Lou handles all biblography operations.

For example:
	- Reformat artciles and their references downloaded by Kopernico.
	- Relocate articles manually renamed.
`,
	Example: `
Lou biblo reformat`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"reformat"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// find home directory
		home, errHomedir := homedir.Dir()
		if errHomedir != nil {
			fmt.Println(errHomedir)
			os.Exit(1)
		}

		switch args[0] {
		case "format":

			// lineBreaks
			aux.LineBreaks()

			// buffers
			var stdout bytes.Buffer
			var stderr bytes.Buffer

			// shellCall
			commd := home + "/Factorem/Lou/sh/format.sh"
			shCmd := exec.Command(commd)

			// run
			shCmd.Stdout = &stdout
			shCmd.Stderr = &stderr
			err := shCmd.Run()

			if err != nil {
				log.Printf(chalk.Red.Color("error: %v\n"), err)
			}

			// stdout
			fmt.Println(chalk.Cyan.Color(stdout.String()))

			// stderr
			if stderr.String() != "" {
				fmt.Println(chalk.Red.Color(stderr.String()))
			}

			// lineBreaks
			aux.LineBreaks()

		}
	},

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

func init() {
	rootCmd.AddCommand(bibloCmd)

	////////////////////////////////////////////////////////////////////////////////////////////////////

	// flags

	////////////////////////////////////////////////////////////////////////////////////////////////////

}
