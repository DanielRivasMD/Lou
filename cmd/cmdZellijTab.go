// cmd/tab.go
/*
Copyright © 2025 Daniel Rivas

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU GPL v3, or (at your option) any later version.
*/
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	tabTarget    string
	validLayouts = map[string]string{
		"tab":     "Default tab layout",
		"explore": "Explore layout",
		"repl":    "REPL layout",
	}
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var tabCmd = &cobra.Command{
	Use:   "tab [layout]",
	Short: "Launch a new Zellij tab",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) +
		chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

Launch a Zellij tab with one of the following layouts:

  tab      - ` + validLayouts["tab"] + `
  explore  - ` + validLayouts["explore"] + `
  repl     - ` + validLayouts["repl"] + `

Provide the layout as an argument (defaults to "tab"), and optionally pass --target.`,
	Example: chalk.Cyan.Color("lou") + " tab explore --target ~/code",

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Args:              cobra.MaximumNArgs(1),
	ValidArgs:         []string{"tab", "explore", "repl"},
	ValidArgsFunction: completeLayoutNames,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		layout := "tab"
		if len(args) == 1 {
			layout = args[0]
		}

		createTab(layout)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(tabCmd)

	tabCmd.Flags().StringVarP(&tabTarget, "target", "t", "", "Change to this directory before launching")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// completeLayoutNames offers tab-completion for the layout argument.
func completeLayoutNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	var out []string
	for name := range validLayouts {
		if toComplete == "" || startsWith(name, toComplete) {
			out = append(out, name)
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}

// helper to check prefix
func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

////////////////////////////////////////////////////////////////////////////////////////////////////
