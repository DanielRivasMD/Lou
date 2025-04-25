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
	hidden bool
	no_ignore bool
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// countCmd
var countCmd = &cobra.Command{
	Use:   "count [dir|file]",
	Aliases: []string{"c"},
	Short:  "Count directories or files in the current location",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` efficiently counts directories or files in the specified target location
Options for hidden data and ignoring configurations are included for flexible usage
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("count") + ` dir` + `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("count") + ` file` + `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("count") + ` --hidden file` + `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Yellow.Color("count") + ` --no-ignore dir` + `
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"dir", "file"},
	Args:      cobra.ExactValidArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// base command
		cmdCount := "fd ."

		// append flags
		if hidden {
			cmdCount += " --hidden"
		}

		// append flags
		if no_ignore {
			cmdCount += " --no-ignore"
		}

		// validate input
		arg := args[0]

		// args
		switch arg {
			case "dir":
				cmdCount += ` --type=d --max-depth=1 | /usr/bin/wc -l`
			case "file":
				cmdCount += ` --type=f --max-depth=1 | /usr/bin/wc -l`
			default:
				fmt.Printf("Invalid argument: %s\n", arg)
		}

		// execute command
		shellCall(cmdCount)
	},

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(countCmd)

	// flags
	countCmd.Flags().BoolVarP(&hidden, "hidden", "H", false, "Account for hidden data")
	countCmd.Flags().BoolVarP(&no_ignore, "no-ignore", "I", false, "Do not respect ignore config")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
