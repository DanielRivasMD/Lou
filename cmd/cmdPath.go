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
	"os"
	"os/exec"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// pathCmd prints one of the standard shell path variables.
// Usage: lou path [path|fpath|gopath]  (defaults to PATH).
var pathCmd = &cobra.Command{
	Use:   "path [which]",
	Short: "Print entries of PATH, FPATH, or GOPATH",
	Long: "Show each element of the chosen colon-delimited shell variable.\n" +
		"If no argument is given, PATH is used.\n" +
		"Valid options: path, fpath, gopath",
	Args:              cobra.MaximumNArgs(1),
	ValidArgs:         []string{"path", "fpath", "gopath"},
	ValidArgsFunction: completePathVars,
	Run: func(cmd *cobra.Command, args []string) {
		var envvar string
		switch {
		case len(args) == 0:
			envvar = "PATH"
		case args[0] == "path" || args[0] == "PATH":
			envvar = "PATH"
		case args[0] == "fpath" || args[0] == "FPATH":
			envvar = "FPATH"
		case args[0] == "gopath" || args[0] == "GOPATH":
			envvar = "GOPATH"
		default:
			fmt.Fprintf(os.Stderr, "invalid option %q, must be one of [path, fpath, gopath]\n", args[0])
			os.Exit(1)
		}
		printEnvPaths(envvar)
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// completePathVars offers shell-completion for the three path types.
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

// printEnvPaths fetches $ENVVAR, splits on ':', and prints each segment on its own line.
func printEnvPaths(envvar string) {
	// 1) read via shell to allow expansion
	cmd := exec.Command("sh", "-c", "echo $"+envvar)
	out, err := cmd.Output()
	horus.CheckErr(err, horus.WithOp("path.print"), horus.WithMessage("reading "+envvar))

	// 2) split and print
	parts := strings.Split(strings.TrimSpace(string(out)), ":")
	domovoi.LineBreaks(true)
	for i, p := range parts {
		// avoid trailing newline at end
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
