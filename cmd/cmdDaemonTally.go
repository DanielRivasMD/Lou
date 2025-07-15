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
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//
// Subcommand: tally
//

var tallyCmd = &cobra.Command{
	Use:     "tally",
	Short:   "List all daemons started by Lou",
	Aliases: []string{"list", "ls"},
	Run: func(cmd *cobra.Command, args []string) {
		dir := getDaemonDir()
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read daemon dir:", err)
			os.Exit(1)
		}

		fmt.Printf("%-20s %-15s %-6s %s\n", "NAME", "GROUP", "PID", "STATUS")
		for _, e := range entries {
			if e.IsDir() {
				continue
			}
			name := e.Name()[:len(e.Name())-len(filepath.Ext(e.Name()))]
			meta, err := loadMeta(name)
			if err != nil {
				continue
			}
			status := chalk.Red.Color("stopped")
			if p, err := os.FindProcess(meta.PID); err == nil {
				if err = p.Signal(syscall.Signal(0)); err == nil {
					status = chalk.Green.Color("running")
				}
			}
			fmt.Printf("%-20s %-15s %-6d %s\n",
				meta.Name, meta.Group, meta.PID, status)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////


func init() {
	daemonCmd.AddCommand(tallyCmd)
}


////////////////////////////////////////////////////////////////////////////////////////////////////
