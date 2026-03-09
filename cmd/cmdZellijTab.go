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
		&rootFlags.tabLayout,
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
		rootFlags.tabTarget = ""
	case 1:
		rootFlags.tabTarget = args[0]
	default:
		horus.CheckErr(
			fmt.Errorf("too many arguments: %d", len(args)),
			horus.WithOp(op),
			horus.WithCategory("USAGE_ERROR"),
			horus.WithMessage("tab takes at most one directory argument"),
		)
	}

	createTab(rootFlags.tabLayout, rootFlags.tabTarget)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
