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
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//
// Subcommand: summon
//

var summonCmd = &cobra.Command{
	Use:   "summon [name]",
	Short: "View logs for a daemon",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		meta, err := loadMeta(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "No such daemon %q\n", name)
			os.Exit(1)
		}
		pager := os.Getenv("PAGER")
		if pager == "" {
			pager = "less"
		}
		c := exec.Command(pager, "-R", meta.LogPath)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		c.Run()
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	daemonCmd.AddCommand(summonCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
