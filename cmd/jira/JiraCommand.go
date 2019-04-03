package jira

import (
	"github.com/bmaximilian/gutils/cmd/jira/token"
	"github.com/bmaximilian/gutils/cmd/jira/worklogs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serverKeyFlag string
var Command = &cobra.Command{
	Use:   "jira",
	Short: "CLI Tools for Jira actions",
}

// Set the default viper values
func SetDefaults() {
	viper.SetDefault("jira.url", "https://jira.atlassian.com")
	viper.SetDefault("jira.apiPrefix", "/rest/api/2")
	viper.SetDefault("jira.cert", nil)
	viper.SetDefault("jira.key", nil)
	viper.SetDefault("jira.token", nil)
	worklogs.SetDefaults()
	token.SetDefaults()
}

// Initializes the command line tool
func InitCommand() {
	Command.PersistentFlags().StringVarP(
		&serverKeyFlag,
		"server",
		"s",
		"",
		"The Server to use",
	)

	worklogs.InitCommand()
	Command.AddCommand(worklogs.Command)

	token.InitCommand()
	Command.AddCommand(token.Command)
}
