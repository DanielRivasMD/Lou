////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// parseShell extracts metadata from function definition
func parseShell(content string) ([]Function, error) {
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
		functions = append(functions, Function{
			Shell:       "zsh",
			Name:        match[re.SubexpIndex("name")],
			Description: strings.TrimSpace(match[re.SubexpIndex("desc")]),
			Arguments:   strings.TrimSpace(match[re.SubexpIndex("args")]),
		})
	}

	return functions, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
