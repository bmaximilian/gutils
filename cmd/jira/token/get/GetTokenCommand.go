package get

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/fatih/color"
	googleLogger "github.com/google/logger"
	"github.com/spf13/cobra"
	"log"
)

var username string
var password string

var Command = &cobra.Command{
	Use:   "get",
	Short: "Get a authorization token from username and password",
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)
		token, err := connect.GetTokenFromUserNameAndPassword(
			cmd.Flag("user").Value.String(),
			cmd.Flag("password").Value.String(),
		)

		if err != nil {
			l.Fatalln(err)
		}

		l.Infof(
			"Jira Token : %v\n",
			color.CyanString("%v", token),
		)
	},
}

// Set the default viper values
func SetDefaults() {}

// Initializes the command line tool
func InitCommand() {
	Command.Flags().StringVarP(
		&username,
		"user",
		"u",
		"",
		"The username for the token",
	)

	Command.Flags().StringVarP(
		&password,
		"password",
		"p",
		"",
		"The password for the token",
	)
}
