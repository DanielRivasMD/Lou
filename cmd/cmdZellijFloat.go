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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var zfCmd = &cobra.Command{}

////////////////////////////////////////////////////////////////////////////////////////////////////

// 1) Define type to hold parameters
type LayoutGeometry struct {
	Width, Height, X, Y string
}

// 2) Declare a map from layout name → geometry
var layoutPresets = map[string]LayoutGeometry{
	"full": {
		Width:  "100%",
		Height: "100%",
		X:      "0",
		Y:      "0",
	},
	"half-left": {
		Width:  "50%",
		Height: "100%",
		X:      "0",
		Y:      "0",
	},
	"half-right": {
		Width:  "50%",
		Height: "100%",
		X:      "50%",
		Y:      "0",
	},
	"top-left": {
		Width:  "45%",
		Height: "45%",
		X:      "0",
		Y:      "0",
	},
	"bottom-left": {
		Width:  "45%",
		Height: "45%",
		X:      "0",
		Y:      "60%",
	},
	"top-right": {
		Width:  "45%",
		Height: "45%",
		X:      "60%",
		Y:      "0",
	},
	"bottom-right": {
		Width:  "45%",
		Height: "45%",
		X:      "60%",
		Y:      "60%",
	},
}

// 3) Derive validLayouts slice from map keys
var validLayouts = func() []string {

	keys := make([]string, 0, len(layoutPresets))
	for name := range layoutPresets {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	return keys
}()

var layoutFlag string

// flags for zellij floats
var (
	floatHeight string
	floatWidth  string
	floatX      string
	floatY      string
)

// mirror the defaults you passed into StringVarP
var (
	defaultHeight = "100%"
	defaultWidth  = "95%"
	defaultX      = "10"
	defaultY      = "0"
)

func init() {
	rootCmd.AddCommand(zfCmd)

	zfCmd.PersistentFlags().StringVarP(&floatHeight, "height", "H", defaultHeight, "pane height as percentage")
	zfCmd.PersistentFlags().StringVarP(&floatWidth, "width", "W", defaultWidth, "pane width as percentage")
	zfCmd.PersistentFlags().StringVarP(&floatX, "x", "X", defaultX, "horizontal offset as percentage")
	zfCmd.PersistentFlags().StringVarP(&floatY, "y", "Y", defaultY, "vertical offset as percentage")

	layoutPresets["default"] = LayoutGeometry{
		Width:  floatWidth,
		Height: floatHeight,
		X:      floatX,
		Y:      floatY,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// pickValue applies bottom -> top override:
//  1. start with defaultVal
//  2. if presetVal is non-empty, use
//  3. if flagVal differs from defaultVal, use
func pickValue(defaultVal, presetVal, flagVal string) string {
	result := defaultVal
	if presetVal != "" {
		result = presetVal
	}
	if flagVal != defaultVal {
		result = flagVal
	}
	return result
}

// look up named preset, apply overrides from flags, and return final LayoutGeometry
func resolveLayoutGeometry(layoutName string) (LayoutGeometry, error) {

	geom, ok := layoutPresets[layoutName]
	if !ok {
		return LayoutGeometry{}, fmt.Errorf(
			"unknown layout %q (must be one of %v)",
			layoutName, validLayouts,
		)
	}

	// override
	geom.Width = pickValue(defaultWidth, geom.Width, floatWidth)
	geom.Height = pickValue(defaultHeight, geom.Height, floatHeight)
	geom.X = pickValue(defaultX, geom.X, floatX)
	geom.Y = pickValue(defaultY, geom.Y, floatY)

	return geom, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
