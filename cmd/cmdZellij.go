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
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	// zellijCmd is the parent for all Zellij-specific subcommands.
	zellijCmd = &cobra.Command{
		Use:   "zellij",
		Short: "Commands to interact with Zellij terminal multiplexer",
		Long:  "Group Zellij-related actions under a single command namespace.",
	}

	// zellijTabCmd is the 'lou zellij tab' subcommand.
	zellijTabCmd = &cobra.Command{
		Use:     "tab",
		Short:   "Launch a new Zellij tab with a custom layout",
		Long:    "Under the zellij namespace, spins up a new tab using your defined layout.",
		Example: "lou zellij tab --target /path/to/dir",
		Run: func(cmd *cobra.Command, args []string) {
			createTab("tab")
		},
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zellijCmd)
	zellijCmd.AddCommand(zellijTabCmd)

	// reuse the same --target flag
	zellijTabCmd.Flags().
		StringVarP(&tabTarget, "target", "t", "", "Change to this directory before launching the tab")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
