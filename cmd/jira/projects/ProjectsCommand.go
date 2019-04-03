package projects

import (
	"github.com/bmaximilian/gutils/cmd/jira/projects/list"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "projects",
	Short: "JIRA Project Actions",
}

// Set the default viper values
func SetDefaults() {
	list.SetDefaults()
}

// Initializes the command line tool
func InitCommand() {
	list.InitCommand()
	Command.AddCommand(list.Command)
}
