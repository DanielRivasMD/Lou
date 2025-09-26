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

var zellijCmd = &cobra.Command{}

// TODO: write readme documentation on customizing tab layouts
var tabCmd = &cobra.Command{
	Use:     "tab [path]",
	Short:   "Launch a new Zellij tab",
	Long:    helpTab,
	Example: exampleTab,

	Args: cobra.MaximumNArgs(1),

	Run: runTab,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(tabCmd)

	tabCmd.Flags().StringVarP(&flags.tabLayout, "layout", "l", flags.tabLayout, "Layout to use [tab, explore, repl]")

	tabCmd.RegisterFlagCompletionFunc("layout",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var out []string
			for name := range validTabLayouts {
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
	const op = "tab.cmd"

	// TODO: define tab layout
	flags.tabLayout = "tab"

	// Validate positional arguments (at most one path)
	switch len(args) {
	case 0:
		flags.tabTarget = ""
	case 1:
		flags.tabTarget = args[0]
	default:
		horus.CheckErr(
			fmt.Errorf("too many arguments: %d", len(args)),
			horus.WithOp(op),
			horus.WithMessage("tab takes at most one directory argument"),
		)
	}

	createTab(flags.tabLayout, flags.tabTarget)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
