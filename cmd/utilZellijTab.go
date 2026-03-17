////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func createTab(tabType, tabTarget string) {
	const op = "lou.tab.create"

	if _, ok := validTabTypes[tabType]; !ok {
		validTypes := make([]string, 0, len(validTabTypes))
		for t := range validTabTypes {
			validTypes = append(validTypes, t)
		}
		horus.CheckErr(
			fmt.Errorf("invalid workspace type %q", tabType),
			horus.WithOp(op),
			horus.WithCategory("VALIDATION_ERROR"),
			horus.WithMessage(fmt.Sprintf("must be one of: %s", strings.Join(validTypes, ", "))),
		)
	}

	cmdStr := fmt.Sprintf(
		`zellij action new-tab --layout $HOME/.lou/layouts/%s.kdl --name $( [ "$PWD" = "$HOME" ] && echo "~" || basename "$PWD" )`,
		tabType,
	)

	if tabTarget == "" {
		if err := domovoi.ExecSh(cmdStr); err != nil {
			horus.CheckErr(
				err,
				horus.WithOp(op),
				horus.WithCategory("ZELLIJ_ERROR"),
				horus.WithMessage("failed to launch tab"),
			)
		}
		return
	}

	orig, err := domovoi.RecallDir()
	if err != nil {
		horus.CheckErr(
			err,
			horus.WithOp(op),
			horus.WithCategory("DIR_ERROR"),
			horus.WithMessage("failed to recall working directory"),
		)
	}

	defer func() {
		if err := domovoi.ChangeDir(orig); err != nil {
			horus.CheckErr(
				err,
				horus.WithOp(op),
				horus.WithCategory("DIR_ERROR"),
				horus.WithMessage("failed to restore original directory (tab was created)"),
				horus.WithExitCode(0),
			)
		}
	}()

	if err := domovoi.ChangeDir(tabTarget); err != nil {
		horus.CheckErr(
			err,
			horus.WithOp(op),
			horus.WithCategory("DIR_ERROR"),
			horus.WithMessage("failed to change to target directory"),
		)
	}

	if err := domovoi.ExecSh(cmdStr); err != nil {
		horus.CheckErr(
			err,
			horus.WithOp(op),
			horus.WithCategory("ZELLIJ_ERROR"),
			horus.WithMessage("failed to launch tab"),
		)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
