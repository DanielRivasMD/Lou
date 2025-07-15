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
	"time"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//
// Subcommand: invoke
//

var (
	daemonName string
	watchDir   string
	scriptPath string
	groupName  string
	logPath    string
)

var invokeCmd = &cobra.Command{
	Use:   "invoke",
	Short: "Start a new watcher daemon",
	Run: func(cmd *cobra.Command, args []string) {
		// validate required flags
		if daemonName == "" {
			fmt.Fprintln(os.Stderr, "Error: --name is required")
			os.Exit(1)
		}
		if watchDir == "" {
			fmt.Fprintln(os.Stderr, "Error: --watch is required")
			os.Exit(1)
		}
		if scriptPath == "" {
			fmt.Fprintln(os.Stderr, "Error: --script is required")
			os.Exit(1)
		}

		// expand env and tildes
		watchDir = os.ExpandEnv(watchDir)
		scriptPath = os.ExpandEnv(scriptPath)
		home := os.Getenv("HOME")

		// default log path
		if logPath == "" {
			logPath = filepath.Join(home, ".lou", "logs", daemonName+".log")
		} else {
			logPath = os.ExpandEnv(logPath)
		}

		// build metadata
		meta := &daemonMeta{
			Name:       daemonName,
			Group:      groupName,
			WatchDir:   watchDir,
			ScriptPath: scriptPath,
			LogPath:    logPath,
			InvokedAt:  time.Now(),
		}

		// spawn watcher
		pid, err := spawnWatcher(meta)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to start watcher:", err)
			os.Exit(1)
		}
		meta.PID = pid

		// persist metadata
		if err := saveMeta(meta); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to write metadata:", err)
			os.Exit(1)
		}

		fmt.Printf("%s invoked daemon %q (group=%q) with PID %d\n",
			chalk.Green.Color("OK:"), daemonName, groupName, pid)
	},
}

func init() {
	daemonCmd.AddCommand(invokeCmd)

	invokeCmd.Flags().StringVarP(&daemonName, "name", "n", "", "Unique daemon name")
	invokeCmd.Flags().StringVarP(&watchDir, "watch", "w", "", "Directory to watch")
	invokeCmd.Flags().StringVarP(&scriptPath, "script", "s", "", "Script to run on change")
	invokeCmd.Flags().StringVarP(&groupName, "group", "g", "", "Daemon group name")
	invokeCmd.Flags().StringVarP(&logPath, "log", "l", "", "Path for log file")

}

////////////////////////////////////////////////////////////////////////////////////////////////////
