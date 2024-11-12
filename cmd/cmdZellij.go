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
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// zellijCmd
var zellijCmd = &cobra.Command{

	Use:   "zellij",
	Short: "" + chalk.Yellow.Color("cobra") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` help ` + chalk.Yellow.Color("zellij"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"control", "explore", "repl", "go", "julia", "rust"},
	Args:                  cobra.ExactValidArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// $HOME/.archive/in-situ/zellij/layouts/*kdl
		cmd := "zellij --layout " + findHome() + "/.archive/in-situ/zellij/layouts/" + args[0] + ".kdl"

		fmt.Println(cmd)
		shcmd := cmd
		ε, _, _ := shellCall(shcmd)
		checkErr(ε)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(zellijCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
