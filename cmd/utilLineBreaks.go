////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	// "log"

	"github.com/labstack/gommon/color"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// print line breaks
func lineBreaks() {
	for ι := 0; ι < 100; ι++ {
		color.Print(color.Grey("=", color.B))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
