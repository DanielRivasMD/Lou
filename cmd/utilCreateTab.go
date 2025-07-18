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

// createTab performs the common logic for launching tabs.
func createTab(layout string) {
	const op = "tab.launch"

	if _, ok := validLayouts[layout]; !ok {
		horus.CheckErr(
			fmt.Errorf("invalid layout %q", layout),
			horus.WithOp(op),
			horus.WithMessage("must be one of [tab, explore, repl]"),
		)
	}

	cmdStr := fmt.Sprintf(
		`zellij action new-tab --layout $HOME/.lou/layouts/%s.kdl --name $( [ "$PWD" = "$HOME" ] && echo "~" || basename "$PWD" )`,
		layout,
	)

	if tabTarget != "" {
		orig, err := domovoi.RecallDir()
		if err != nil {
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to recall working directory"))
		}
		if err := domovoi.ChangeDir(tabTarget); err != nil {
			domovoi.ChangeDir(orig)
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to change to target directory"))
		}
		if err := domovoi.ExecSh(cmdStr); err != nil {
			domovoi.ChangeDir(orig)
			horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to launch tab"))
		}
		return
	}

	if err := domovoi.ExecSh(cmdStr); err != nil {
		horus.CheckErr(err, horus.WithOp(op), horus.WithMessage("failed to launch tab"))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
