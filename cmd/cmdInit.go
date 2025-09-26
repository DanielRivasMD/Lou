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
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "",
	Long:    helpInit,
	Example: exampleInit,

	Run: runInit,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(initCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runInit(cmd *cobra.Command, args []string) {
	const op = "lou.init"

	// name each for nicer error messages
	toCreate := []struct {
		label, path string
	}{
		{"lou root", dirs.lou},
		{"layout", dirs.layout},
		{"shell", dirs.sh},
	}

	for _, dir := range toCreate {
		horus.CheckErr(
			domovoi.CreateDir(dir.path, flags.verbose),
			horus.WithOp(op),
			horus.WithCategory("io_error"),
			horus.WithMessage(fmt.Sprintf("creating %s directory", dir.label)),
			horus.WithDetails(map[string]any{
				"path": dir.path,
			}),
		)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
