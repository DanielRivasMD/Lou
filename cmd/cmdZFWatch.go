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

	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var zWatchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "launch watcher on a floating zellij window with ease",
	Long:    helpZFWatch,
	Example: exampleZFWatch,

	Run: runWatch,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zWatchCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runWatch(cmd *cobra.Command, args []string) {
	// TODO: test whether `just watch` exists
	// TODO: can persistent flags be overwritten?
	cmdFloat := fmt.Sprintf(`
		zellij run --name watch --close-on-exit --floating --pinned true \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, floatHeight, floatWidth, floatX, floatY)
	cmdFloat += `just watch`
	domovoi.ExecCmd("bash", "-c", cmdFloat)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
