////////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
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

var helpZFBat = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"View data in a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window using "+chalk.Cyan.Color(chalk.Italic.TextStyle("bat"))+" for a specified file",
)

var helpZFEza = domovoi.FormatHelp(
	"Daniel Rivas",
	"<danielrivasmd@gmail.com>",
	"View data in a floating "+chalk.Cyan.Color(chalk.Italic.TextStyle("zellij"))+" window using "+chalk.Cyan.Color(chalk.Italic.TextStyle("eza")),
)

var helpZFFloat = domovoi.FormatHelp(
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
