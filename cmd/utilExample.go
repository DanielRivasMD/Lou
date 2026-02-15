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

var exampleGit = domovoi.FormatExample(
	"lou",
	[]string{"git"},
)

var exampleKnit = domovoi.FormatExample(
	"lou",
	[]string{"knit"},
)

var exampleMap = domovoi.FormatExample(
	"lou",
	[]string{"map"},
)

var exampleShow = domovoi.FormatExample(
	"lou",
	[]string{"show"},
)

var exampleCount = domovoi.FormatExample(
	"lou",
	[]string{"diff"},
)

var exampleDiff = domovoi.FormatExample(
	"lou",
	[]string{"diff"},
)

var exampleDoctor = domovoi.FormatExample(
	"lou",
	[]string{"doctor"},
)

var exampleHelper = domovoi.FormatExample(
	"lou",
	[]string{"helper"},
)

var exampleHidden = domovoi.FormatExample(
	"lou",
	[]string{"hidden"},
)

var examplePath = domovoi.FormatExample(
	"lou",
	[]string{"path"},
)

var exampleFloat = domovoi.FormatExample(
	"lou",
	[]string{"float"},
)

var exampleLazygit = domovoi.FormatExample(
	"lou",
	[]string{"lazygit"},
)

var exampleBat = domovoi.FormatExample(
	"lou",
	[]string{"bat", "<file>"},
)

var exampleEza = domovoi.FormatExample(
	"lou",
	[]string{"eza", "<path>"},
)

var exampleTab = domovoi.FormatExample(
	"lou",
	[]string{" tab", "~/src/helix", "--layout", "explore"},
)

var exampleResize = domovoi.FormatExample(
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

// TODO: update editor example
// // helpEditorExample returns the example usage snippet.
// func helpEditorExample(editor string) string {
// 	return chalk.White.Color("lou") + " " +
// 		chalk.White.Color(chalk.Bold.TextStyle(editor)) + " " +
// 		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>"))
// }

////////////////////////////////////////////////////////////////////////////////////////////////////
