package worklogs

import (
	"github.com/bmaximilian/gutils/cmd/jira/worklogs/list"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "worklogs",
	Short: "CRUD Jira Worklogs",
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
