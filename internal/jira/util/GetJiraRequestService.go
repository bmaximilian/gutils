package util

import (
	"errors"
	"fmt"
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/spf13/viper"
	"reflect"
)

func getValidServerConfigKey(serverConfigKey string) (string, error) {
	servers := viper.GetStringMap("jira.servers")
	serverKeys := reflect.ValueOf(servers).MapKeys()

	if len(serverKeys) == 0 {
		return "", errors.New("no jira server configuration found in config file")
	}

	if serverConfigKey == "" {
		return serverKeys[0].String(), nil
	} else {
		for serverName := range servers {
			if serverName == serverConfigKey {
				return serverName, nil
			}
		}
	}

	return "", errors.New(fmt.Sprintf("no jira server configuration found in config file for key %v", serverConfigKey))
}

func NewJiraServerConfigFromViper(viperKey string) *connect.JiraServerConfig {
	config := &connect.JiraServerConfig{
		APIVersion: 2,
		Tempo:      &connect.TempoOptions{},
		TlsConfig:  &connect.TLSConfig{},
	}

	if value := viper.Get(viperKey + ".url"); value != nil {
		config.Url = value.(string)
	}
	if value := viper.Get(viperKey + ".apiVersion"); value != nil {
		config.APIVersion = value.(int)
	}
	if value := viper.Get(viperKey + ".user"); value != nil {
		config.UserName = value.(string)
	}
	if value := viper.Get(viperKey + ".password"); value != nil {
		config.Password = value.(string)
	}
	if value := viper.Get(viperKey + ".tls.cert"); value != nil {
		config.TlsConfig.CertPath = value.(string)
	}
	if value := viper.Get(viperKey + ".tls.key"); value != nil {
		config.TlsConfig.KeyPath = value.(string)
	}
	if value := viper.Get(viperKey + ".tls.password"); value != nil {
		config.TlsConfig.Password = value.(string)
	}
	if value := viper.Get(viperKey + ".token"); value != nil {
		config.Token = value.(string)
	}
	if value := viper.Get(viperKey + ".tempo"); value != nil && value.(bool) {
		config.Tempo.Enabled = true
		config.Tempo.ApiVersion = 2
	}

	if config.UserName != "" && config.Password != "" && config.Token == "" {
		var err error = nil
		config.Token, err = connect.GetTokenFromUserNameAndPassword(config.UserName, config.Password)

		if err != nil {
			l := logger.GetLogger()
			l.Warningln(err)
		}
	}

	return config
}

func GetJiraRequestService(serverConfigKey string) *connect.JiraRequestService {
	l := logger.GetLogger()

	validServerConfigKey, getKeyErr := getValidServerConfigKey(serverConfigKey)
	if getKeyErr != nil {
		l.Fatalln(getKeyErr)
	}

	serverConfigString := "jira.servers." + validServerConfigKey

	jiraRequestService, tokenParseError := connect.NewJiraRequestService(
		NewJiraServerConfigFromViper(serverConfigString),
	)

	if tokenParseError != nil {
		l.Fatalln(tokenParseError)
	}

	return jiraRequestService
}
