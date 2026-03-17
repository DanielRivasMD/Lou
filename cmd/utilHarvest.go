/*
Copyright © 2026 Daniel Rivas <danielrivasmd@gmail.com>

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
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func copyToClipboard(data []byte) error {
	switch runtime.GOOS {
	case "darwin":
		cmd := exec.Command("pbcopy")
		cmd.Stdin = bytes.NewReader(data)
		return cmd.Run()
	case "linux":
		if _, err := exec.LookPath("copyq"); err == nil {
			cmd := exec.Command("copyq", "copy", "-")
			cmd.Stdin = bytes.NewReader(data)
			return cmd.Run()
		}
		return fmt.Errorf("no clipboard utility found: install copyq")
	default:
		return fmt.Errorf("unsupported platform for clipboard: %s", runtime.GOOS)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
