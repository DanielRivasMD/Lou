/*
Copyright © 2025 Daniel Rivas <danielrivasmd@gmail.com>

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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var helperCmd = &cobra.Command{
	Use:     "helper",
	Short:   "" + chalk.Yellow.Color("") + ".",
	Long:    helpHelper,
	Example: exampleHelper,

	Run: runHelper,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(helperCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runHelper(cmd *cobra.Command, args []string) {

	arg := args[0]

	cmdHelper := `
# grab the list under “Available Commands:”
cmds=( $(
  ` + arg + ` help | 
  sed -n '/Available Commands:/,/Flags:/p' |
  grep -E '^[[:space:]]+[a-z]' | 
  awk '{print $1}'
) )

for sub in "${cmds[@]}"; do
  printf "==================================================\n"
  printf "===== %s \n" "$sub"
  printf "==================================================\n"
  ` + arg + ` help "$sub"
done
`

	domovoi.LineBreaks(true)
	domovoi.ExecCmd("zsh", "-c", cmdHelper)
	domovoi.LineBreaks(true)

}

////////////////////////////////////////////////////////////////////////////////////////////////////
