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
	"time"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//
// Subcommand: rekindle
//

var rekindleCmd = &cobra.Command{
	Use:   "rekindle [name]",
	Short: "Restart a stopped daemon",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		meta, err := loadMeta(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "No such daemon %q\n", name)
			os.Exit(1)
		}
		pid, err := spawnWatcher(meta)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to restart:", err)
			os.Exit(1)
		}
		meta.PID = pid
		meta.InvokedAt = time.Now()
		if err := saveMeta(meta); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to update metadata:", err)
			os.Exit(1)
		}
		fmt.Printf("%s rekindled %q with PID %d\n",
			chalk.Green.Color("OK:"), name, pid)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	daemonCmd.AddCommand(rekindleCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
