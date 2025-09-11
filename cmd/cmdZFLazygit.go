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

var zfLazygitCmd = &cobra.Command{
	Use:     "lazygit",
	Aliases: []string{"lg"},
	Short:   "lazygit in a floating zellij window",
	Long:    helpZFLazygit,
	Example: exampleZFLazygit,

	Run: runLazygit,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zfLazygitCmd)
	zfCmd.AddCommand(zfLazygitCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runLazygit(cmd *cobra.Command, args []string) {

	layoutName := "default"
	if len(args) == 1 {
		layoutName = args[0]
	}

	geom, _ := resolveLayoutGeometry(layoutName)

	cmdLazygit := fmt.Sprintf(`
		zellij run --name lazygit --close-on-exit --floating --pinned true \
		--height %s \
		--width %s \
		--x %s \
		--y %s \
		-- `, geom.Height, geom.Width, geom.X, geom.Y)
	cmdLazygit += `lazygit`
	domovoi.ExecCmd("bash", "-c", cmdLazygit)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
