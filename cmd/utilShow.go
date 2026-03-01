////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func calculatePadding(functions []Function) (shellPad, namePad, descPad, argsPad int) {
	shellPad, namePad, descPad, argsPad = len("Shell"), len("Function"), len("Description"), len("Arguments")

	for _, fn := range functions {
		shellPad = max(shellPad, len(fn.Shell))
		namePad = max(namePad, len(fn.Name)+2)         // +2 for backticks
		descPad = max(descPad, len(fn.Description))
		argsPad = max(argsPad, len(fn.Arguments)+2)    // +2 for backticks
	}
	return
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func parseFile(path string) ([]Function, error) {
	var allFuncs []Function
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	funcs, err := parseShell(string(content))
	if err != nil {
		return nil, err
	}
	allFuncs = append(allFuncs, funcs...)
	return allFuncs, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func parseShell(content string) ([]Function, error) {
	var functions []Function

	re := regexp.MustCompile(
		`#\s*function:\s*(?P<name>\w+).*?\n` +
			`#\s*description:\s*(?P<desc>[^\n]+).*?\n` +
			`(#\s*arguments:\s*(?P<args>[^\n]+).*?\n)?` +
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

func generateMD(functions []Function) string {
	var builder strings.Builder
	shellPad, namePad, descPad, argsPad := calculatePadding(functions)

	builder.WriteString("# Shell Functions Documentation\n\n")
	builder.WriteString(fmt.Sprintf(
		"| %-*s | %-*s | %-*s | %-*s |\n",
		shellPad, "Shell",
		namePad, "Function",
		descPad, "Description",
		argsPad, "Arguments",
	))

	builder.WriteString(fmt.Sprintf(
		"|%s|%s|%s|%s|\n",
		strings.Repeat("-", shellPad+2),
		strings.Repeat("-", namePad+2),
		strings.Repeat("-", descPad+2),
		strings.Repeat("-", argsPad+2),
	))

	for _, fn := range functions {
		argDisplay := fmt.Sprintf("`%s`", fn.Arguments)
		if fn.Arguments == "" {
			argDisplay = "``"
		}

		builder.WriteString(fmt.Sprintf(
			"| %-*s | %-*s | %-*s | %-*s |\n",
			shellPad, fn.Shell,
			namePad, fmt.Sprintf("`%s`", fn.Name),
			descPad, fn.Description,
			argsPad, argDisplay,
		))
	}
	return builder.String()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
