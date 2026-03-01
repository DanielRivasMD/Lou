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
	"os/exec"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	pathCmd := MakeCmd("path", runPath,
		WithArgs(cobra.MaximumNArgs(1)),
		WithValidArgs([]string{"path", "fpath", "gopath"}),
	)
	pathCmd.ValidArgsFunction = completePathVars
	rootCmd.AddCommand(pathCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runPath(cmd *cobra.Command, args []string) {
	const op = "lou.path"

	var envvar string
	switch {
	case len(args) == 0:
		envvar = "PATH"
	case strings.EqualFold(args[0], "path"):
		envvar = "PATH"
	case strings.EqualFold(args[0], "fpath"):
		envvar = "FPATH"
	case strings.EqualFold(args[0], "gopath"):
		envvar = "GOPATH"
	default:
		horus.CheckErr(
			fmt.Errorf("invalid option %q, must be one of [path, fpath, gopath]", args[0]),
			horus.WithOp(op),
			horus.WithCategory("USAGE_ERROR"),
			horus.WithExitCode(1),
		)
	}
	printEnvPaths(envvar)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func completePathVars(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	choices := []string{"path", "fpath", "gopath"}
	var out []string
	for _, c := range choices {
		if strings.HasPrefix(c, strings.ToLower(toComplete)) {
			out = append(out, c)
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func printEnvPaths(envvar string) {
	const op = "lou.path.print"

	cmd := exec.Command("sh", "-c", "echo $"+envvar)
	out, err := cmd.Output()
	horus.CheckErr(
		err,
		horus.WithOp(op),
		horus.WithMessage("reading "+envvar),
		horus.WithCategory("SYS_CMD"),
	)

	parts := strings.Split(strings.TrimSpace(string(out)), ":")
	domovoi.LineBreaks(true)
	for i, p := range parts {
		if i == len(parts)-1 {
			fmt.Print(p)
		} else {
			fmt.Println(p)
		}
	}
	fmt.Println()
	domovoi.LineBreaks(true)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
