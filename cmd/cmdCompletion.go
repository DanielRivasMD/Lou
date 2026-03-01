/*
Copyright © 2021 Daniel Rivas <danielrivasmd@gmail.com>

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
	"os"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	completionCmd := MakeCmd("completion", runCompletion,
		WithArgs(cobra.ExactArgs(1)),
		WithValidArgs([]string{"bash", "zsh", "fish", "powershell"}),
	)
	completionCmd.DisableFlagsInUseLine = true // Can't set via options yet
	rootCmd.AddCommand(completionCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCompletion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		horus.CheckErr(cmd.Root().GenBashCompletion(os.Stdout))
	case "zsh":
		horus.CheckErr(cmd.Root().GenZshCompletion(os.Stdout))
	case "fish":
		horus.CheckErr(cmd.Root().GenFishCompletion(os.Stdout, true))
	case "powershell":
		horus.CheckErr(cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
