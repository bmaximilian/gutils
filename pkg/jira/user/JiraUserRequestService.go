package user

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/jira/user/models"
	"github.com/levigross/grequests"
)

type JiraUserRequestService struct {
	requestService *connect.JiraRequestService
}

func NewJiraUserRequestService(jiraRequestService *connect.JiraRequestService) *JiraUserRequestService {
	return &JiraUserRequestService{requestService: jiraRequestService}
}

func (j *JiraUserRequestService) GetUser(username string, token string) (*models.User, error) {
	parsedResponse := models.User{}
	response, err := j.requestService.Get("/user", &grequests.RequestOptions{
		Params: map[string]string{
			"username": username,
		},
		Headers: map[string]string{
			"Authorization": "Basic " + token,
		},
	})
	if err != nil {
		return nil, err
	}

	jsonErr := response.JSON(&parsedResponse)
	if jsonErr != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
