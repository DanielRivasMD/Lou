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

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	path string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// zellijViewCmd
var zellijViewCmd = &cobra.Command{

	Use:   "zview",
	Short: "" + chalk.Yellow.Color("cobra") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` help ` + chalk.Yellow.Color("zellij"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"bat", "hx", "micro", "lsd"},
	Args:      cobra.ExactValidArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////


	Run: func(κ *cobra.Command, args []string) {

		// base command
		cmdView := `zellij run --name canvas --floating --height 100 --width 100 --x 100 --y 0 -- `

		// validate input
		arg := args[0]

		switch arg {
			case "bat":
				cmdView += `bat `
			case "hx":
				cmdView += `hx `
			case "micro":
				cmdView += `micro `
			case "lsd":
				cmdView += `lsd  --header --long --classify --git `
			default:
				fmt.Printf("Invalid argument: %s\n", arg)
		}

		cmdView += path

		shellCall(cmdView)

		// cmdZellij := findHome() + "/" + ".lou" + "/" + "sh" + "/" + "zellij.sh"
		// execCmd(cmdZellij, fn.String(), path)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zellijViewCmd)

	// flags
	zellijViewCmd.Flags().StringVarP(&path, "path", "p", "", "Data path")

	zellijViewCmd.MarkFlagRequired("path")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
