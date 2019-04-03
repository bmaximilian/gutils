package connect

import (
	"encoding/base64"
	"errors"
)

// Create a jira authorization token for passed username and password
func GetTokenFromUserNameAndPassword(username string, password string) (string, error) {
	if username == "" {
		return "", errors.New("username can not be empty")
	}

	if password == "" {
		return "", errors.New("password can not be empty")
	}

	return base64.StdEncoding.EncodeToString([]byte(username + ":" + password)), nil
}
