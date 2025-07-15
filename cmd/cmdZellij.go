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

// tabTarget is the directory to cd into before creating a tab.
var (
	tabTarget string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	// zellijCmd is the parent for all Zellij-specific subcommands.
	zellijCmd = &cobra.Command{
		Use:   "zellij",
		Short: "Commands to interact with Zellij terminal multiplexer",
		Long:  "Group Zellij-related actions under a single command namespace.",
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zellijCmd)

}

////////////////////////////////////////////////////////////////////////////////////////////////////
