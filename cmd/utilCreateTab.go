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
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// createTab performs the common logic for both commands.
func createTab(layoutType string) {
	const op = "cmd.createTab"

	// Allowed layouts
	valids := map[string]bool{"tab": true, "explore": true, "repl": true}
	if !valids[layoutType] {
		horus.CheckErr(
			fmt.Errorf("invalid layout %q", layoutType),
			horus.WithOp(op),
			horus.WithMessage("must be one of [tab, explore, repl]"),
		)
	}

	// Build the Zellij action
	cmdStr := fmt.Sprintf(
		`zellij action new-tab --layout $HOME/.lou/layouts/%s.kdl --name $( [ $PWD = $HOME ] && echo "~" || basename $PWD )`,
		layoutType,
	)

	// Optionally change directory
	if tabTarget != "" {
		orig, err := domovoi.RecallDir()
		if err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to recall directory"))
		}
		if err := domovoi.ChangeDir(tabTarget); err != nil {
			domovoi.ChangeDir(orig)
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to change directory"))
		}
		if err := domovoi.ExecSh(cmdStr); err != nil {
			domovoi.ChangeDir(orig)
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to launch new tab"))
		}
		return
	}

	// No directory switch
	if err := domovoi.ExecSh(cmdStr); err != nil {
		horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to launch new tab"))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
