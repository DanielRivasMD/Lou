////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/labstack/gommon/color"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func lineBreaks() {
	for ι := 0; ι < 80; ι++ {
		color.Print(color.Grey("=", color.B))
	}

	ç := 1
	for ç <= 2 {
		fmt.Println()
		ç++
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
