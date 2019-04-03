package util

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/spf13/viper"
)

func GetJiraRequestService() *connect.JiraRequestService {
	l := logger.GetLogger()

	jiraRequestService, tokenParseError := connect.NewJiraRequestService(
		viper.Get("jira.url").(string),
		viper.Get("jira.apiPrefix").(string),
		viper.Get("jira.token").(string),
		viper.Get("jira.cert").(string),
		viper.Get("jira.key").(string),
	)

	if tokenParseError != nil {
		l.Fatalln(tokenParseError)
	}

	return jiraRequestService
}
