package get

import (
	"fmt"
	"github.com/bmaximilian/gutils/internal/jira/util"
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/jira/user"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/fatih/color"
	googleLogger "github.com/google/logger"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"syscall"
)

var username string
var password string

var Command = &cobra.Command{
	Use:   "get",
	Short: "Get a authorization token from username and password",
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)

		server := cmd.Flag("server").Value.String()
		pwd := cmd.Flag("password").Value.String()
		usr := cmd.Flag("user").Value.String()

		if pwd == "" {
			l.Info("Please enter your JIRA Password: ")
			bytePwd, readErr := terminal.ReadPassword(int(syscall.Stdin))
			if readErr != nil {
				l.Fatalln(readErr)
			}

			pwd = string(bytePwd)
			fmt.Println()
		}

		token, err := connect.GetTokenFromUserNameAndPassword(
			usr,
			pwd,
		)

		if err != nil {
			l.Fatalln(err)
		}

		jiraUserRequestService := user.NewJiraUserRequestService(util.GetJiraRequestService(server))

		res, getUserErr := jiraUserRequestService.GetUser(usr, token)
		if getUserErr != nil {
			l.Fatalln(getUserErr)
		}

		if res == nil {
			l.Warningln("could not fetch user for credentials. Credentials may be wrong")
		} else if res.DisplayName != "" {
			l.Infof("User for Token: %v\n", res.DisplayName)
		}

		fmt.Println()
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
