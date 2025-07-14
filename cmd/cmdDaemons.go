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
	"encoding/json"
	"fmt"
	// "io/fs"
	"os"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

type DaemonsMeta struct {
	Name    string `json:"name"`
	Group   string `json:"group"`
	PID     int    `json:"pid"`
	LogPath string `json:"logPath"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func getDaemonDir() string {
	return filepath.Join(os.Getenv("HOME"), ".lou", "daemons")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var daemonsCmd = &cobra.Command{
	Use:   "daemons",
	Short: "List all daemons started by Lou",

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		dir := getDaemonDir()
		entries, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println(chalk.Red.Color("Failed to read daemon directory:"), err)
			os.Exit(1)
		}

		fmt.Printf("%-20s %-15s %-6s %s\n", "NAME", "GROUP", "PID", "STATUS")
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			full := filepath.Join(dir, entry.Name())
			data, err := os.ReadFile(full)
			if err != nil {
				continue
			}
			var d DaemonsMeta
			if err := json.Unmarshal(data, &d); err != nil {
				continue
			}

			// Check if process is alive
			status := chalk.Red.Color("stopped")
			if proc, err := os.FindProcess(d.PID); err == nil {
				if err = proc.Signal(syscall.Signal(0)); err == nil {
					status = chalk.Green.Color("running")
				}
			}

			fmt.Printf(
				"%-20s %-15s %-6d %s\n",
				d.Name,
				d.Group,
				d.PID,
				status,
			)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(daemonsCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
