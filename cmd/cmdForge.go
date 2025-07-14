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

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	watchDir   string
	scriptPath string
	groupName  string
	logPath    string
	restart    bool
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// forgeCmd
var forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "Watch a directory and rebuild your Karabiner config on changes",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
		`

Lou’s forge command spins up a background watcher that triggers your build script
whenever files in a given directory change. Logs, auto-restart, and grouping supported.
`,
	Example: chalk.Cyan.Color("lou") +
		" forge --watch ~/.archive/in-situ/karabiner/frag \\\n" +
		"     --script ~/.archive/in-silico/.forge/goku.sh \\\n" +
		"     --group karabiner --log ~/.lou/logs/forge.log\n",

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		const op = "cmd.forge"

		// 1) Validate inputs
		if watchDir == "" {
			horus.CheckErr(
				fmt.Errorf("`--watch` is required"),
				horus.WithOp(op),
				horus.WithMessage("no directory to watch specified"),
			)
		}

		if scriptPath == "" {
			horus.CheckErr(
				fmt.Errorf("`--script` is required"),
				horus.WithOp(op),
				horus.WithMessage("no script to run specified"),
			)
		}

		// Expand ~
		watchDir = os.ExpandEnv(watchDir)
		scriptPath = os.ExpandEnv(scriptPath)

		// 2) Prepare watcher command (using watchexec)
		//    You can install watchexec via Homebrew: brew install watchexec
		cmdArgs := []string{
			"--watch", watchDir,
			"--",
			"bash", "-c", fmt.Sprintf("source %s", scriptPath),
		}
		watcher := exec.Command("watchexec", cmdArgs...)

		// 3) Redirect logs if requested
		if logPath != "" {
			// ensure log directory exists
			logDir := filepath.Dir(logPath)
			if err := os.MkdirAll(logDir, 0755); err != nil {
				horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to create log directory"))
			}
			f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to open log file"))
			}
			watcher.Stdout = f
			watcher.Stderr = f
		}

		// 4) Start detached
		if err := watcher.Start(); err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to start watcher"))
		}

		// 5) Inform the user and exit
		fmt.Printf(
			"%s watcher started (group=%q) with PID %d\n",
			chalk.Green.Color("OK:"), groupName, watcher.Process.Pid,
		)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(forgeCmd)

	// Bind flags to Viper defaults
	forgeCmd.Flags().StringVarP(&watchDir, "watch", "w", viper.GetString("watch"), "Directory to watch for changes")
	forgeCmd.Flags().StringVarP(&scriptPath, "script", "s", viper.GetString("script"), "Script to execute on file change")
	forgeCmd.Flags().StringVarP(&groupName, "group", "g", viper.GetString("group"), "Daemon group name")
	forgeCmd.Flags().StringVarP(&logPath, "log", "l", viper.GetString("log"), "Path to log file")
	forgeCmd.Flags().BoolVar(&restart, "restart", viper.GetBool("restart"), "Automatically restart the watcher on failure")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
