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
	"os"
	"strings"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

// Function represents a parsed shell function
type Function struct {
	Shell       string
	Name        string
	Description string
	Arguments   string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// listCmd
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("") + ` help ` + chalk.Yellow.Color("shell") + chalk.Yellow.Color("list"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// collect documentation
		functions, err := parseFile(inFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing functions: %v\n", err)
			os.Exit(1)
		}

		// generate & print Markdown
		markdown := generateMarkdown(functions)
		fmt.Println(markdown)

	},

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// parseShellFunction extracts metadata from function definition
func parseShellFunction(content string) ([]Function, error) {
	var functions []Function

	// regex match zsh/bash function descriptions
	// example format:
	//   # function: fn
	//   # description: cool function
	//   # arguments: some args
	//   fn() { ... }
	re := regexp.MustCompile(
		`#\s*function:\s*(?P<name>\w+).*?\n` +
		`#\s*description:\s*(?P<desc>[^\n]+).*?\n` +
		`(#\s*arguments:\s*(?P<args>[^\n]+).*?\n)?` +  // Optional arguments line
		`(?P<code>(?P<name2>\w+)\(\)\s*\{[^}]+\})`,
	)

	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		name := match[re.SubexpIndex("name")]
		desc := match[re.SubexpIndex("desc")]
		args := match[re.SubexpIndex("args")]

		functions = append(functions, Function{
			Shell:       "zsh",
			Name:        name,
			Description: strings.TrimSpace(desc),
			Arguments: args,
		})
	}

	return functions, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// generateMarkdown creates a Markdown table from parsed functions
func generateMarkdown(functions []Function) string {
	var builder strings.Builder

	builder.WriteString("# Shell Functions Documentation\n\n")
	builder.WriteString("| Shell | Function | Description | Arguments |\n")
	builder.WriteString("|-------|----------|-------------|-----------|\n")

	for _, fn := range functions {
		builder.WriteString(fmt.Sprintf(
			"| %s | `%s` | %s | `%s` |\n",
			fn.Shell, fn.Name, fn.Description, fn.Arguments,
		))
	}

	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func parseFile(path string) ([]Function, error) {
	var allFuncs []Function
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	funcs, err := parseShellFunction(string(content))
	if err != nil {
		return nil, err
	}
	allFuncs = append(allFuncs, funcs...)
	return allFuncs, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	shellCmd.AddCommand(listCmd)

	// flags
	listCmd.MarkFlagRequired("file")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
