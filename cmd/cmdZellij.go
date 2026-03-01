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
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var validTabTypes = map[string]string{
	"devel":   "Default development layout",
	"tab":     "Single vanilla pane",
	"tabs2":   "Two stacked vanilla panes",
	"tabs3":   "Three stacked vanilla panes",
	"tabs4":   "Four stacked vanilla panes",
	"tabs5":   "Five stacked vanilla panes",
	"explore": "Two stacked panes: top runs `y`, bottom vanilla",
	"repl":    "Editor + canvas + right-side repl pane",
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	tabCmd := MakeCmd("tab", runTab,
		WithArgs(cobra.MaximumNArgs(1)),
	)
	rootCmd.AddCommand(tabCmd)

	tabCmd.Flags().StringVarP(
		&flags.tabLayout,
		"type",
		"t",
		"devel",
		"Workspace type: devel, tab, tabs2, tabs3, tabs4, tabs5, explore, repl",
	)

	// Shell completion
	tabCmd.RegisterFlagCompletionFunc("type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var out []string
			for name := range validTabTypes {
				if toComplete == "" || strings.HasPrefix(name, toComplete) {
					out = append(out, name)
				}
			}
			return out, cobra.ShellCompDirectiveNoFileComp
		},
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runTab(cmd *cobra.Command, args []string) {
	const op = "lou.tab"

	switch len(args) {
	case 0:
		flags.tabTarget = ""
	case 1:
		flags.tabTarget = args[0]
	default:
		horus.CheckErr(
			fmt.Errorf("too many arguments: %d", len(args)),
			horus.WithOp(op),
			horus.WithCategory("USAGE_ERROR"),
			horus.WithMessage("tab takes at most one directory argument"),
		)
	}

	createTab(flags.tabLayout, flags.tabTarget)
}

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

	if tabTarget != "" {
		orig, err := domovoi.RecallDir()
		if err != nil {
			horus.CheckErr(
				err,
				horus.WithOp(op),
				horus.WithCategory("DIR_ERROR"),
				horus.WithMessage("failed to recall working directory"),
			)
		}
		if err := domovoi.ChangeDir(tabTarget); err != nil {
			domovoi.ChangeDir(orig)
			horus.CheckErr(
				err,
				horus.WithOp(op),
				horus.WithCategory("DIR_ERROR"),
				horus.WithMessage("failed to change to target directory"),
			)
		}
		if err := domovoi.ExecSh(cmdStr); err != nil {
			domovoi.ChangeDir(orig)
			horus.CheckErr(
				err,
				horus.WithOp(op),
				horus.WithCategory("ZELLIJ_ERROR"),
				horus.WithMessage("failed to launch tab"),
			)
		}
		return
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
