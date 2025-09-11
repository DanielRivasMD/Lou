////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var helpRoot = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"",
)

var helpInit = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"",
)

var helpBat = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"View data in a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window using "+chalk.Cyan.Color(chalk.Italic.TextStyle("bat"))+" for a specified file",
)

var helpEza = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"View data in a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window using "+chalk.Cyan.Color(chalk.Italic.TextStyle("eza")),
)

var helpFloat = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"Launch a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window using "+chalk.Cyan.Color(chalk.Italic.TextStyle("float"))+" with ease",
)

var helpZFLazygit = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"lazygit in a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window using "+chalk.Cyan.Color(chalk.Italic.TextStyle("lazygit")),
)

var helpZFResize = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"Resize one random floating pane to a percentage of screen size and move it to an anchor point",
)

var helpZFWatch = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"launch a "+chalk.Cyan.Color(chalk.Italic.TextStyle("watcher"))+" on a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window with ease",
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: update edirot helps
// helpEditorShort returns the one‐line Short description.
func helpEditorShort(editor string) string {
	return fmt.Sprintf("view data in a floating zellij window using %s", editor)
}

// helpEditorLong returns the multi‐line Long description.
func helpEditorLong(editor string) string {
	header := chalk.Green.Color(
		chalk.Bold.TextStyle("Daniel Rivas "),
	) +
		chalk.Dim.TextStyle(
			chalk.Italic.TextStyle("<danielrivasmd@gmail.com>"),
		)

	body := fmt.Sprintf(
		"\n\nview data in a floating %szellij%s window using %s",
		chalk.Cyan.Color(""),
		chalk.Cyan.Color(""),
		chalk.Cyan.Color(editor),
	)

	return header + chalk.Dim.TextStyle(body)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
