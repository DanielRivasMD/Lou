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
	"os"
	"os/exec"
	"path/filepath"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	// "github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var summonCmd = &cobra.Command{
	Use:   "summon [name]",
	Short: "Open the log file for a Lou daemon",
	Args:  cobra.ExactArgs(1),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		const op = "cmd.summon"
		name := args[0]
		metaFile := filepath.Join(getDaemonDir(), name+".json")

		data, err := os.ReadFile(metaFile)
		horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to read metadata"))

		var d DaemonsMeta
		horus.CheckErr(json.Unmarshal(data, &d), horus.WithOp(op), horus.WithMessage("invalid metadata"))

		if d.LogPath == "" {
			horus.CheckErr(
				fmt.Errorf("no log path defined for %q", name),
				horus.WithOp(op),
				horus.WithMessage("cannot summon logs"),
			)
		}

		fmt.Println(d.LogPath)

		// Launch pager
		pager := exec.Command("bat", "--paging", "always", d.LogPath)
		pager.Stdin = os.Stdin
		pager.Stdout = os.Stdout
		pager.Stderr = os.Stderr

		horus.CheckErr(pager.Run(), horus.WithOp(op), horus.WithMessage("error running pager"))
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(summonCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
