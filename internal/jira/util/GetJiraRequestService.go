package util

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/spf13/viper"
)

func GetJiraRequestService() *connect.JiraRequestService {
	l := logger.GetLogger()

	certP := viper.Get("jira.cert")
	keyP := viper.Get("jira.key")
	tlsConfig := connect.TLSConfig{}

	if certP != nil {
		tlsConfig.CertPath = certP.(string)
	}

	if keyP != nil {
		tlsConfig.KeyPath = keyP.(string)
	}

	jiraRequestService, tokenParseError := connect.NewJiraRequestService(
		viper.Get("jira.url").(string),
		viper.Get("jira.apiPrefix").(string),
		viper.Get("jira.token").(string),
		&tlsConfig,
	)

	if tokenParseError != nil {
		l.Fatalln(tokenParseError)
	}

	return jiraRequestService
}
