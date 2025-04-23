////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// generateMD creates Markdown table from parsed functions
func generateMD(functions []Function) string {
	var builder strings.Builder
	shellPad, namePad, descPad, argsPad := calculatePadding(functions)

	// header
	builder.WriteString("# Shell Functions Documentation\n\n")
	builder.WriteString(fmt.Sprintf(
		"| %-*s | %-*s | %-*s | %-*s |\n",
		shellPad, "Shell",
		namePad, "Function",
		descPad, "Description",
		argsPad, "Arguments",
	))

	// separator
	builder.WriteString(fmt.Sprintf(
		"|%s|%s|%s|%s|\n",
		strings.Repeat("-", shellPad+2),
		strings.Repeat("-", namePad+2),
		strings.Repeat("-", descPad+2),
		strings.Repeat("-", argsPad+2),
	))

	// rows
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


