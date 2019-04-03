package token

import (
	"github.com/bmaximilian/gutils/cmd/jira/token/get"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "token",
	Short: "Do something to get a JIRA Authorization token",
}

// Set the default viper values
func SetDefaults() {
	get.SetDefaults()
}

// Initializes the command line tool
func InitCommand() {
	get.InitCommand()
	Command.AddCommand(get.Command)
}
