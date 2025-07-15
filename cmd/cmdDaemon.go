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
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

//
// Metadata structure and helpers
//

// daemonMeta holds persistent info about a watcher process.
type daemonMeta struct {
	Name       string    `json:"name"`
	Group      string    `json:"group"`
	WatchDir   string    `json:"watchDir"`
	ScriptPath string    `json:"scriptPath"`
	LogPath    string    `json:"logPath"`
	PID        int       `json:"pid"`
	InvokedAt  time.Time `json:"invokedAt"`
}

// getDaemonDir returns ~/.lou/daemons
func getDaemonDir() string {
	return filepath.Join(os.Getenv("HOME"), ".lou", "daemons")
}

// saveMeta writes meta to ~/.lou/daemons/<name>.json
func saveMeta(meta *daemonMeta) error {
	dir := getDaemonDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return err
	}
	path := filepath.Join(dir, meta.Name+".json")
	return os.WriteFile(path, data, 0644)
}

// loadMeta reads ~/.lou/daemons/<name>.json
func loadMeta(name string) (*daemonMeta, error) {
	path := filepath.Join(getDaemonDir(), name+".json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var meta daemonMeta
	err = json.Unmarshal(data, &meta)
	return &meta, err
}

// spawnWatcher starts watchexec, redirects logs, returns its PID
func spawnWatcher(meta *daemonMeta) (int, error) {
	// ensure log directory exists
	if err := os.MkdirAll(filepath.Dir(meta.LogPath), 0755); err != nil {
		return 0, err
	}

	cmdArgs := []string{
		"--watch", meta.WatchDir,
		"--",
		"bash", meta.ScriptPath,
	}
	cmd := exec.Command("watchexec", cmdArgs...)

	// open (or create) log file
	f, err := os.OpenFile(meta.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	cmd.Stdout = f
	cmd.Stderr = f

	// start detached
	if err := cmd.Start(); err != nil {
		f.Close()
		return 0, err
	}
	pid := cmd.Process.Pid

	// detach to avoid zombies
	if err := cmd.Process.Release(); err != nil {
		return pid, err
	}
	return pid, nil
}

//
// Parent Command: lou daemon
//

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage background watcher daemons",
	Long: chalk.Green.Color("lou daemon") + ` is a namespace for controlling
background watcher processes.  Subcommands:

  invoke    start a new watcher
  tally     list all daemons
  slay      stop a daemon by name
  summon    view logs for a daemon
  rekindle  restart a stopped daemon`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(daemonCmd)
	// daemonCmd.AddCommand(invokeCmd, tallyCmd, slayCmd, summonCmd, rekindleCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
