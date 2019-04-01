package cmd

import (
	"fmt"
	"github.com/bmaximilian/gutils/cmd/calculate"
	"github.com/bmaximilian/gutils/cmd/jira"
	"github.com/bmaximilian/gutils/cmd/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{Use: "gutils"}

// Set the default viper values
func SetDefaults() {
	calculate.SetDefaults()
	jira.SetDefaults()
}

// Initializes the command line tool
func init() {
	rootCmd.AddCommand(version.VersionCommand)

	calculate.InitCalculateCommand()
	rootCmd.AddCommand(calculate.CalculateCommand)

	jira.InitCommand()
	rootCmd.AddCommand(jira.Command)
}

// Executes the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
