// //////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"os"

	"github.com/DanielRivasMD/horus"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func matchDir(directoryPath string) {
	// Open the target directory
	dir, err := os.Open(directoryPath)
	horus.CheckErr(err)
	defer dir.Close()

	// Read all entries in the directory
	entries, err := dir.Readdir(0)
	horus.CheckErr(err)

	// Flag to track if at least one file was removed.
	removedAny := false

	// Check each file @ the directory
	for _, entry := range entries {
		if fileNamePattern.MatchString(entry.Name()) {
			removedAny = true
			fullPath := directoryPath + entry.Name()
			fmt.Println(fullPath)
			err = os.Remove(fullPath)
			horus.CheckErr(err)
			fmt.Printf("Removed file: %s\n", fullPath)
		}
	}

	// Trigger if no matching files were found to remove
	if !removedAny {
		fmt.Println(chalk.Cyan.Color("\tNo files to remove"))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
