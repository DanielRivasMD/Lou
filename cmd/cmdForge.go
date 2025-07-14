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
	"encoding/json"
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
	daemonName string
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
	Short: "Master of all daemons",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
		`

Lou’s forge command spins up a background watcher that triggers your build script
whenever files in a given directory change. Logs, auto-restart, and grouping supported.
`,
	Example: chalk.Cyan.Color("lou") +
		" forge --name goku-forge \\\n" +
		"     --watch ~/.archive/in-situ/karabiner/frag \\\n" +
		"     --script ~/.archive/in-silico/.forge/goku.sh \\\n" +
		"     --group karabiner --log ~/.lou/logs/forge.log\n",

	// TODO: segregate as it own cli tool => lilith
	// TODO: add `watch` as an option

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		const op = "cmd.forge"

		// 1) Validate required inputs
		if daemonName == "" {
			horus.CheckErr(
				fmt.Errorf("`--name` is required"),
				horus.WithOp(op),
				horus.WithMessage("no daemon name specified"),
			)
		}
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

		// 2) Expand environment variables and tildes
		watchDir = os.ExpandEnv(watchDir)
		scriptPath = os.ExpandEnv(scriptPath)

		// 3) Build the watcher command (watchexec)
		cmdArgs := []string{
			"--watch", watchDir,
			"--",
			"bash", scriptPath,
		}
		watcher := exec.Command("watchexec", cmdArgs...)

		// 4) Redirect logs if requested
		if logPath != "" {
			logPath = os.ExpandEnv(logPath)
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

		// 5) Start watcher in background
		if err := watcher.Start(); err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to start watcher"))
		}

		pid := watcher.Process.Pid

		// 6) Write metadata JSON for management commands
		daemonDir := filepath.Join(os.Getenv("HOME"), ".lou", "daemons")
		if err := os.MkdirAll(daemonDir, 0755); err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to create daemon metadata directory"))
		}

		meta := struct {
			Name    string `json:"name"`
			Group   string `json:"group"`
			PID     int    `json:"pid"`
			LogPath string `json:"logPath"`
		}{
			Name:    daemonName,
			Group:   groupName,
			PID:     pid,
			LogPath: logPath,
		}

		data, err := json.MarshalIndent(meta, "", "  ")
		if err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to marshal daemon metadata"))
		}

		metafile := filepath.Join(daemonDir, daemonName+".json")
		if err := os.WriteFile(metafile, data, 0644); err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to write daemon metadata file"))
		}

		// 7) Inform the user and exit
		fmt.Printf(
			"%s daemon %q started (group=%q) with PID %d\n",
			chalk.Green.Color("OK:"), daemonName, groupName, pid,
		)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(forgeCmd)

	// Bind flags to Viper defaults
	forgeCmd.Flags().StringVarP(&daemonName, "name", "n", viper.GetString("name"), "Unique name for this daemon (required)")
	forgeCmd.Flags().StringVarP(&watchDir, "watch", "w", viper.GetString("watch"), "Directory to watch for changes")
	forgeCmd.Flags().StringVarP(&scriptPath, "script", "s", viper.GetString("script"), "Script to execute on file change")
	forgeCmd.Flags().StringVarP(&groupName, "group", "g", viper.GetString("group"), "Daemon group name")
	forgeCmd.Flags().StringVarP(&logPath, "log", "l", viper.GetString("log"), "Path to log file")
	forgeCmd.Flags().BoolVar(&restart, "restart", viper.GetBool("restart"), "Automatically restart the watcher on failure")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
