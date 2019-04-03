package projects

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/jira/projects/models"
	"github.com/levigross/grequests"
)

type JiraProjectsRequestService struct {
	requestService *connect.JiraRequestService
}

func NewJiraProjectsRequestService(jiraRequestService *connect.JiraRequestService) *JiraProjectsRequestService {
	return &JiraProjectsRequestService{requestService: jiraRequestService}
}

func (j *JiraProjectsRequestService) GetAllProjects() (*[]models.Project, error) {
	parsedResponse := make([]models.Project, 0)
	response, err := j.requestService.Get("/project", &grequests.RequestOptions{})
	if err != nil {
		return nil, err
	}

	jsonErr := response.JSON(&parsedResponse)
	if jsonErr != nil {
		return nil, err
	}

	return &parsedResponse, nil
}
