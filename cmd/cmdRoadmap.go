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

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
	"os"
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var roadmapCmd = &cobra.Command{
	Use:   "roadmap",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("") + ` help ` + chalk.Yellow.Color("") + chalk.Yellow.Color("roadmap"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		path := args[0]
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", path, err)
			os.Exit(1)
		}

		text := string(data)

		// split on lines containing only '=' characters
		re := regexp.MustCompile(`(?m)^[=]+\s*$`)
		rawBlocks := re.Split(text, -1)

		seen := make(map[string]struct{})
		var unique []string
		sepLine := "=================================================="

		for _, blk := range rawBlocks {
			blk = strings.TrimSpace(blk)
			if blk == "" {
				continue
			}
			if _, exists := seen[blk]; !exists {
				seen[blk] = struct{}{}
				unique = append(unique, blk)
			}
		}

		// reassemble with separator lines
		result := strings.Join(unique, "\n"+sepLine+"\n")
		fmt.Println(result)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(roadmapCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
