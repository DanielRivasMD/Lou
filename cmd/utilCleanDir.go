////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func cleanDir(location string) {
	// Insert visual separators to enhance console readability.
	domovoi.LineBreaks()

	// Invoke matchDir to remove matching files from the specified directory.
	matchDir(location)

	// Initialize buffers for capturing any process output.
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Display any captured standard output in cyan.
	fmt.Println(chalk.Cyan.Color(stdout.String()))

	// If any error output is captured, display it in red.
	if stderr.String() != "" {
		fmt.Println(chalk.Red.Color(stderr.String()))
	}

	// Insert additional visual separators to conclude the output block.
	domovoi.LineBreaks()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
