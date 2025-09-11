////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/domovoi"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var exampleRoot = domovoi.FormatExample(
	"lou",
	[]string{"help"},
)

var exampleInit = domovoi.FormatExample(
	"lou",
	[]string{"init"},
)

var exampleZFFloat = domovoi.FormatExample(
	"lou",
	[]string{"float"},
)

var exampleZFLazygit = domovoi.FormatExample(
	"lou",
	[]string{"lazygit"},
)

var exampleZFBat = domovoi.FormatExample(
	"lou",
	[]string{"bat", "<file>"},
)

var exampleZFEza = domovoi.FormatExample(
	"lou",
	[]string{"eza", "<path>"},
)

var exampleZFResize = domovoi.FormatExample(
	"lou",
	[]string{"resize"},
	[]string{"resize", "--height", "<100%>", "--width", "<95%>", "--x", "<10>", "--y", "<0>"},
	[]string{"resize", "full"},
	[]string{"resize", "half-left"},
	[]string{"resize", "half-right"},
)

var exampleZFWatch = domovoi.FormatExample(
	"lou",
	[]string{"watch"},
)

////////////////////////////////////////////////////////////////////////////////////////////////////
